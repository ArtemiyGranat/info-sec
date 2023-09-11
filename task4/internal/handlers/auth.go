package handlers

import (
	"info-sec-api/internal/storage"
	"info-sec-api/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthHandler(db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("usr")
		password := c.PostForm("passwd")

		user, err := storage.AuthUser(db, username)
		// TODO: How should I handle errors other than mongo.ErrNoDocuments?
		if err != nil {
			log.Printf("Could not find user with given username: %v", err)
			c.String(http.StatusForbidden, "Incorrect username or password")
			return
		}

		err = utils.VerifyPassword(user, password)
		if err != nil {
			c.String(http.StatusForbidden, "Incorrect username or password")
			return
		}

		c.String(http.StatusOK, "Correct password")
	}
}