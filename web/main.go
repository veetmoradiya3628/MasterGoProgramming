package main

import (
	"database/sql"
	"log"
	"os"
)

// application holds the dependencies for our web application, such as loggers and the user repository.
type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	userRepo    UserRepository
	templateDir string
	publicPath  string
	tp          *TemplateRenderer
}

func main() {
	db, err := connectToDatabase("users_database.db")
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		errorLog:    log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile),
		infoLog:     log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile),
		userRepo:    NewSQLUserRepository(db),
		templateDir: "./templates",
		publicPath:  "./public",
	}
	app.tp = NewTemplateRenderer(app.templateDir, false) // 2nd parameter isDev is for running in localdevf

	log.Println("Listening on :8080")
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}

// connectToDatabase establishes a connection to the SQLite database and returns the database handle.
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
