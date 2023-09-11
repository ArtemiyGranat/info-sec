package handlers

import (
	"info-sec-api/internal/models"
	storage "info-sec-api/internal/storage/auth"
	crypt "info-sec-api/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func RegistrateHandler(db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("usr")
		password := c.Query("passwd")
	
		salt, err := crypt.GenerateSalt()
		if err != nil {
			log.Fatalf("Could not generate a salt for a password: %v", err)
		}
		hashedPassword, err := crypt.HashPassword(password, salt)
		if err != nil {
			log.Fatalf("Could not hash a password: %v", err)
		}
	
		user := models.User { Username: username, Salt: salt, Password: hashedPassword }
		if err := storage.RegistrateUser(db, user); err != nil {
			c.String(http.StatusBadRequest, "Could not register a user: %v", err)
		} else {
			c.String(http.StatusOK, "User has been registered successfully")
		}
	
	}
}