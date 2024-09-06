package database

import (
	"database/sql"
	"log"
)

func CreateDataTables(db *sql.DB) {

    dropTablesSQL := `
    
    DROP TABLE IF EXISTS Users;
    DROP TABLE IF EXISTS Categories;
    DROP TABLE IF EXISTS Posts;
    DROP TABLE IF EXISTS Comments;
    DROP TABLE IF EXISTS PostLikes;
    DROP TABLE IF EXISTS CommentLikes;
    DROP TABLE IF EXISTS PostDislikes;
    DROP TABLE IF EXISTS CommentDislikes; 
    `
    _, err := db.Exec(dropTablesSQL)
    if err != nil {
        log.Fatalf("Erreur lors de la suppression des tables existantes: %v", err)
    }


	createTablesSQL := `

	    CREATE TABLE IF NOT EXISTS Users (
        user_id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_name TEXT NOT NULL UNIQUE,
        user_email TEXT NOT NULL UNIQUE,
        user_password TEXT NOT NULL
    );
    
    CREATE TABLE IF NOT EXISTS Categories (
        category_id INTEGER PRIMARY KEY AUTOINCREMENT,
        category_name TEXT NOT NULL UNIQUE,
		CHECK(LENGTH(category_name <= 150))
    );

    CREATE TABLE IF NOT EXISTS Posts (
        post_id INTEGER PRIMARY KEY AUTOINCREMENT,
        post_date TEXT NOT NULL,
        post_content TEXT NOT NULL,
        user_id INTEGER,
        category_id INTEGER,
        FOREIGN KEY (user_id) REFERENCES Users(user_id),
        FOREIGN KEY (category_id) REFERENCES Categories(category_id)
    );

    CREATE TABLE IF NOT EXISTS Comments (
        comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
        comment_content TEXT NOT NULL,
        post_id INTEGER,
        post_date TEXT NOT NULL,
        FOREIGN KEY (post_id) REFERENCES Posts(post_id)
    );

    -- Likes et Dislikes pour les posts
    CREATE TABLE IF NOT EXISTS PostLikes (
        like_id INTEGER PRIMARY KEY AUTOINCREMENT,
        post_id INTEGER,
        user_id INTEGER,
        UNIQUE (user_id, post_id),
        FOREIGN KEY (user_id) REFERENCES Users(user_id),
        FOREIGN KEY (post_id) REFERENCES Posts(post_id)
    );

    CREATE TABLE IF NOT EXISTS PostDislikes (
        dislike_id INTEGER PRIMARY KEY AUTOINCREMENT,
        post_id INTEGER,
        user_id INTEGER,
        UNIQUE (user_id, post_id),
        FOREIGN KEY (user_id) REFERENCES Users(user_id),
        FOREIGN KEY (post_id) REFERENCES Posts(post_id)
    );

    -- Likes et Dislikes pour les commentaires
    CREATE TABLE IF NOT EXISTS CommentLikes (
        like_id INTEGER PRIMARY KEY AUTOINCREMENT,
        comment_id INTEGER,
        user_id INTEGER,
        UNIQUE (user_id, comment_id),
        FOREIGN KEY (user_id) REFERENCES Users(user_id),
        FOREIGN KEY (comment_id) REFERENCES Comments(comment_id)
    );

    CREATE TABLE IF NOT EXISTS CommentDislikes (
        dislike_id INTEGER PRIMARY KEY AUTOINCREMENT,
        comment_id INTEGER,
        user_id INTEGER,
        UNIQUE (user_id, comment_id),
        FOREIGN KEY (user_id) REFERENCES Users(user_id),
        FOREIGN KEY (comment_id) REFERENCES Comments(comment_id)
    );
 `   
	_, err = db.Exec(createTablesSQL)
	if err != nil {
		log.Fatalf("Erreur lors de la création des tables: %v", err)
	}

	log.Println("Tables créées avec succès")
}