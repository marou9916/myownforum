package webserver

import (
	"myownforum/internal/database"
	"myownforum/internal/models"
)

// Fonction pour récupérer les commentaires pour un post spécifique
func GetCommentsForPost(postID int) ([]models.Comment, error) {
	db := database.ConnectToDataBase()
	defer db.Close()

	rows, err := db.Query("SELECT comment_id, comment_content, post_id, post_date FROM Comments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.CommentID, &comment.CommentContent, &comment.PostID, &comment.PostDate); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
