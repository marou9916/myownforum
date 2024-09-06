package database

import (
	"golang.org/x/crypto/bcrypt"
)

// Hacher les mots de passe avant de les insérer avec les autres données dans les tables
func HashPassWord(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
