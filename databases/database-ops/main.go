package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"` // Do not expose hashed password in JSON
	CreatedAt      time.Time `json:"created_at"`
}

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	hashed_password TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`

func main() {
	dbName := "data.db"
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

	lastId, err := createUser(db, "Alice", "alice@gmail.com", "hashedpassword1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Created user with ID:", lastId)

	lastId, err = createUser(db, "Bob", "bob@gmail.com", "hashedpassword2")
	if err != nil {
		panic(err)
	}
	fmt.Println("Created user with ID:", lastId)

	lastId, err = createUser(db, "Charlie", "charlie@gmail.com", "hashedpassword3")
	if err != nil {
		panic(err)
	}
	fmt.Println("Created user with ID:", lastId)

	alice, err := GetUserByName(db, "Alice")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Retrieved user: %+v\n", alice)
	}

	user1, err := GetUserByEmail(db, "bob@gmail.com")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Retrieved user: %+v\n", user1)
	}

	user2, err := GetUserByEmail(db, "vm@gmail.com")
	if err != nil {
		fmt.Println("Error retrieving user:", err)
	} else {
		fmt.Printf("Retrieved user: %+v\n", user2)
	}

	users, err := GetUsers(db)
	bs, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("All users: %s\n", bs)
	}
}

func createTable(db *sql.DB, schema string) error {
	_, err := db.Exec(schema)
	return err
}

func createUser(db *sql.DB, name, email, hashedPassword string) (int64, error) {
	stmt := `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`
	hashedPasswordBytes, err := hashPassword(hashedPassword)
	if err != nil {
		return 0, err
	}
	res, err := db.Exec(stmt, name, email, hashedPasswordBytes)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?`
	row := db.QueryRow(stmt, email)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByName(db *sql.DB, name string) (*User, error) {
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users WHERE name = ?`
	row := db.QueryRow(stmt, name)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
