package webserver

import (
    "net/http"
)

//Servir la page d'accueil 
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r,  "./web/index.html")
}