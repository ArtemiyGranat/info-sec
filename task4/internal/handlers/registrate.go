package handlers

import (
	"info-sec-api/internal/models"
	"info-sec-api/internal/storage"
	"info-sec-api/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func RegistrateHandler(db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("usr")
		password := c.PostForm("passwd")
	
		salt, err := utils.GenerateSalt()
		if err != nil {
			log.Fatalf("Could not generate a salt for a password: %v", err)
		}
		hashedPassword, err := utils.HashPassword(password, salt)
		if err != nil {
			log.Fatalf("Could not hash a password: %v", err)
		}
	
		user := models.User { Username: username, Salt: salt, Password: hashedPassword }
		if err := storage.RegistrateUser(db, user); err != nil {
			c.String(http.StatusBadRequest, "Could not register a user")
			return
		} 
	
		c.String(http.StatusOK, "User has been registered successfully")	
	}
}