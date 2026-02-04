package main

// Database transactions
// --------------------------------------------------------

// 1. User creates an account
// 2. Create a wallet for the user
// 3. Want to top up the wallet for the user
// 4. You want to write a transaction log
import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	hashed_password TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`

var profileSchema = `
CREATE TABLE IF NOT EXISTS profiles (
	user_id INTEGER PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
	avatar TEXT,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"` // Do not expose hashed password in JSON
	CreatedAt      time.Time `json:"created_at"`
	Profile        Profile   `json:"profile"`
}

type Profile struct {
	UserID    int       `json:"user_id"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}

// prepared statement helps to optimize repeated queries

func main() {
	dbName := "userdata.db"
	// very important you should not do this in production
	_ = os.Remove(dbName)
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database successfully")

	err = createTable(db, schema)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table created successfully")

	err = createTable(db, profileSchema)
	if err != nil {
		panic(err)
	}
	fmt.Println("Profile table created successfully")

	userID, err := createUser(db, "Bob", "bob@gmail.com", "hashed_password_456", "avatar1.png")
	if err != nil {
		panic(err)
	}
	fmt.Printf("User created successfully with ID: %d\n", userID)

	ctx := context.Background()
	userID2, err := createUserWithContext(ctx, db, "Alice", "alice@gmail.com", "hashed_password_123", "avatar2.png")
	if err != nil {
		panic(err)
	}
	fmt.Printf("User created successfully with ID: %d\n", userID2)

	rowUser, err := GetUserByEmailWithProfile(ctx, db, "alice@gmail.com")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
		} else {
			panic(err)
		}
	}
	user, err := json.MarshalIndent(rowUser, "", "  ")
	fmt.Printf("Retrieved user: %+v\n", string(user))
	user1, err := GetUserByEmailWithProfile(ctx, db, "veet@gmail.com")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
		} else {
			panic(err)
		}
	} else {
		fmt.Printf("Retrieved user: %+v\n", user1)
	}

}

func createTable(db *sql.DB, schema string) error {
	_, err := db.Exec(schema)
	return err
}

// Begin, Commit, Rollback
func createUser(db *sql.DB, name, email, hashedPassword, avatar string) (int64, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	userStmt, err := tx.Prepare("INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer userStmt.Close()

	res, err := userStmt.Exec(name, email, hashedPassword)
	if err != nil {
		return 0, err
	}
	userID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	profileStmt, err := tx.Prepare("INSERT INTO profiles (user_id, avatar) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer profileStmt.Close()

	_, err = profileStmt.Exec(userID, avatar)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func createUserWithContext(ctx context.Context, db *sql.DB, name, email, hashedPassword, avatar string) (int64, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	userStmt, err := tx.PrepareContext(ctx, "INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer userStmt.Close()
	res, err := userStmt.ExecContext(ctx, name, email, hashedPassword)
	if err != nil {
		return 0, err
	}
	userID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	profileStmt, err := tx.PrepareContext(ctx, "INSERT INTO profiles (user_id, avatar) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer profileStmt.Close()
	_, err = profileStmt.ExecContext(ctx, userID, avatar)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func GetUserByEmailWithProfile(ctx context.Context, db *sql.DB, email string) (*User, error) {
	query := `
	SELECT u.id, u.name, u.email, u.hashed_password, u.created_at, p.user_id, p.avatar, p.created_at
	FROM users u
	LEFT JOIN profiles p ON u.id = p.user_id
	WHERE u.email = ?`

	row := db.QueryRowContext(ctx, query, email)
	var user User
	var profile Profile
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt,
		&profile.UserID, &profile.Avatar, &profile.CreatedAt)
	if err != nil {
		return nil, err
	}
	user.Profile = profile
	return &user, nil
}
