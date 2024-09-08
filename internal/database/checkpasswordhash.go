package database

import (
	"golang.org/x/crypto/bcrypt"
)
//Vérifier que le mot de passe rentré par un User correspond au haché dans la base de données
func CheckPassWordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}