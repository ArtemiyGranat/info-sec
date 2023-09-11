package handlers

import (
	"info-sec-api/internal/storage"
	"info-sec-api/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthHandler(db *mongo.Database, address string) gin.HandlerFunc {
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
		
		accessTokenString := utils.NewAccessToken(username)
		refreshTokenString := utils.NewRefreshToken(username)
		c.SetCookie("Access-Token", accessTokenString, int(utils.AccessTokenTTL.Seconds()), "/", address, true, true)
		c.SetCookie("Refresh-Token", refreshTokenString, int(utils.RefreshTokenTTL.Seconds()), "/", address, true, true)

		c.String(http.StatusOK, "Authentification successful")
	}
}