package database

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

// Établir une connexion à la base de données SQLite.
func ConnectToDataBase() *sql.DB {
    db, err := sql.Open("sqlite3", "/home/student/myownforum/migrations/database.sqlite")
    if err != nil {
        log.Fatalf("Erreur lors de la connexion à la base de données: %v", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("Erreur de ping à la base de données: %v", err)
    }

    log.Println("Connexion à la base de données réussie")

    return db
}