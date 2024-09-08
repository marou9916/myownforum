package webserver

import (
	"log"
	"net/http"
)

// Démarrer le serveur HTTP
func StartServer() {
	http.HandleFunc("/", HomeHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}

	log.Println("Serveur démarré sur le port 8080: http://localhost:8080/ .")
}
