package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistrateHandler(c *gin.Context) {
	usr := c.Query("usr")
	passwd := c.Query("passwd")

	c.String(http.StatusOK, "Username is %s and password is %s", usr, passwd)
}