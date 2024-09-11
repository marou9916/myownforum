package webserver

import (
	"html/template"
	"myownforum/internal/database"
	"myownforum/internal/models"
	"net/http"
)

// Servir la page d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//Se connecter à la base de données
	db := database.ConnectToDataBase()
	defer db.Close()

	//Récupérer les données de la base
	selectDataSQL := `SELECT p.post_id, p.post_date, p.post_content, c.category_name 
		FROM Posts p
		JOIN Categories c ON p.category_id = c.category_id
		ORDER BY post_date DESC LIMIT 10`

	rows, err := db.Query(selectDataSQL)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des données", http.StatusInternalServerError)
		return
	}

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err := rows.Scan(&post.PostID, &post.PostDate, &post.PostContent, &post.CategoryName); err != nil {
			http.Error(w, "Erreur lors de la lecture des posts", http.StatusInternalServerError)
			return
		}

		posts = append(posts, post)
	}

	//Charger le template de la page d'accueil
	tmpl, err := template.ParseFiles("/home/student/myownforum/web/templates/index.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}
	//EXécuter le template en y injectant les données
	data := struct {
		Posts []models.Post
	}{
		Posts: posts,
	}

	tmpl.Execute(w, data)
}
