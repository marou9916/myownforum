package database

import (
	"database/sql"
	"log"
)

// Insérer les données dans les tables
func InsertDataIntoTheTable(db *sql.DB) {

	// Hachage des mots de passe
	password1, err := HashPassWord("password123")
	if err != nil {
		log.Fatalf("Erreur lors du hachage du mot de passe : %v", err)
	}

	password2, err := HashPassWord("password456")
	if err != nil {
		log.Fatalf("Erreur lors du hachage du mot de passe : %v", err)
	}

	insertDataSQL := `

		--Insertion des users
	INSERT INTO Users(user_name, user_email, user_password) VALUES
		('JohnDoe', 'johndoe@example.com', '` + password1 + `'),
		('JaneDoe', 'janedoe@example.com', '` + password2 + `');

		--Insertion des users
	INSERT INTO Categories (category_name) VALUES 
		('Technology'),
		('Science'),
		('Lifestyle');

		-- Insertion des posts
		INSERT INTO Posts (post_date, post_content, user_id, category_id) VALUES 
		('2024-09-01', 'First post on technology!', 1, 1),
		('2024-09-02', 'Science is amazing.', 2, 2);

		-- Insertion des commentaires
		INSERT INTO Comments (comment_content, post_id, post_date) VALUES 
		('I totally agree with you!', 1, '2024-09-02'),
		('This is a great post.', 2, '2024-09-03');
	`
	_, err = db.Exec(insertDataSQL)
	if err != nil {
		log.Fatalf("Erreur lors de l'insertion des données: %v", err)
	}

	log.Println("Données insérées avec succès")
}
