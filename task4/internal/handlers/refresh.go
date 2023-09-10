package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefreshHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}