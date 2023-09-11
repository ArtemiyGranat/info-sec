package crypt

import (
	"crypto/rand"
	"info-sec-api/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func GenerateSalt() ([]byte, error) {
    salt := make([]byte, 16)
    _, err := rand.Read(salt)
    if err != nil {
        return nil, err
    }

    return salt, nil
}

func HashPassword(password string, salt []byte) ([]byte, error) {
    passwordBytes := []byte(password + string(salt))
    hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    return hashedPassword, nil
}

func VerifyPassword(user *models.User, password string) error {
    return bcrypt.CompareHashAndPassword(user.Password, []byte(password + string(user.Salt)))
}
