package main

import (
	"fmt"
	"myownforum/internal/database"
)

func main() {
	//Connexion à la base de données
	db := database.ConnectToDataBase()
	defer db.Close()

	//Création des tables
	database.CreateDataTables(db)

	//Insertion des données
	database.InsertDataIntoTheTable(db)

	fmt.Println("Application démarrée")
}
