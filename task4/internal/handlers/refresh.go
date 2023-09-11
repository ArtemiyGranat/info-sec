package handlers

import (
	"info-sec-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RefreshHandler(db *mongo.Database, address string) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshTokenString, err := c.Cookie("Refresh-Token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.String(http.StatusForbidden, "Refresh token not found")
				return
			}
			c.String(http.StatusBadRequest, "Bad request")
			return
		}
		refreshToken, err := utils.ValidateToken(refreshTokenString)
		if err != nil {
			c.String(http.StatusForbidden, "Invalid refresh token")
			return
		}
		username, err := refreshToken.Claims.GetSubject()
		if err != nil {
			c.String(http.StatusForbidden, "Invalid claims in refresh token")
			return
		}

		accessTokenString := utils.NewAccessToken(username)
		newRefreshTokenString := utils.NewRefreshToken(username)
		c.SetCookie("Access-Token", accessTokenString, int(utils.AccessTokenTTL.Seconds()), "/", address, true, true)
		c.SetCookie("Refresh-Token", newRefreshTokenString, int(utils.RefreshTokenTTL.Seconds()), "/", address, true, true)

		c.String(http.StatusOK, "Successful refresh")
	}
}