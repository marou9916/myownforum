package webserver

import (
	"html/template"
	"myownforum/internal/database"
	"myownforum/internal/models"
	"net/http"
)

// Servir la page d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    db := database.ConnectToDataBase()
    defer db.Close()

    rows, err := db.Query("SELECT post_id, post_date, post_content, category_id FROM Posts")
    if err != nil {
        http.Error(w, "Erreur lors de la récupération des posts", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

	var postWithComments []models.PostWithComments
    for rows.Next() {
		//Récupérer les posts
        var post models.Post
        if err := rows.Scan(&post.PostID, &post.PostDate, &post.PostContent, &post.CategoryName); err != nil {
            http.Error(w, "Erreur lors de la lecture des posts", http.StatusInternalServerError)
            return
        }
        
        // Récupérer les commentaires
        comments, err := GetCommentsForPost(post.PostID)
        if err != nil {
            http.Error(w, "Erreur lors de la récupération des commentaires", http.StatusInternalServerError)
            return
        }
        
        postWithComments = append(postWithComments, models.PostWithComments{
			Post: post,
			Comments: comments, 
		})
    }

    tmpl, err := template.ParseFiles("web/templates/index.html")
    if err != nil {
        http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
        return
    }

    data := struct {
        Posts []models.PostWithComments
    }{
        postWithComments,
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, "Erreur lors du rendu du template", http.StatusInternalServerError)
        return
    }

}