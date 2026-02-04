package main

import (
	"database/sql"
	"databases/repository-patten/repository"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbName := "userdata.db"
	db, err := connectToDatabase(dbName)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	fmt.Println("Successfully connected to the database.")

	repo := repository.NewSQLUserRepository(db)
	user, err := repo.GetUsers()
	if err != nil {
		fmt.Printf("Failed to get users: %v\n", err)
		os.Exit(1)
	}
	for _, u := range user {
		output, err := json.MarshalIndent(u, "", "  ")
		if err != nil {
			fmt.Printf("Failed to marshal user: %v\n", err)
			continue
		}
		fmt.Println(string(output))
	}
}

func connectToDatabase(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
