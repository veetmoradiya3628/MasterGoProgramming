package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
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

	userID, err := createUserWithPreparedStmt(db, "Alice", "alice@gmail.com", "hashed_password_123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("User created successfully with ID: %d\n", userID)

	ctx := context.Background()
	userID2, err := createUserwithCtx(ctx, db, "Bob", "bob@alice@gmail.com", "hashed_password_456")
	if err != nil {
		panic(err)
	}
	fmt.Printf("User created successfully with ID: %d\n", userID2)
}

func createTable(db *sql.DB, schema string) error {
	_, err := db.Exec(schema)
	return err
}

func createUserWithPreparedStmt(db *sql.DB, name, email, hashedPassword string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(name, email, hashedPassword)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func createUserwithCtx(ctx context.Context, db *sql.DB, name, email, hashedPassword string) (int64, error) {
	stmt, err := db.PrepareContext(ctx, "INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, name, email, hashedPassword)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
