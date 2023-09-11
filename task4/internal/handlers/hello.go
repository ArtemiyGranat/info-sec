package handlers

import (
	"info-sec-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	accessTokenString, err := c.Cookie("Access-Token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.String(http.StatusForbidden, "Access token not found")
			return
		}
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	_, err = utils.ValidateToken(accessTokenString)
	if err != nil {
		c.String(http.StatusForbidden, "Invalid access token")
	}

	c.String(http.StatusOK, "Hello, InfoSec!")	
}