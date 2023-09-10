package routes

import (
	"info-sec-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func configureRoutes(router *gin.Engine) {
	router.POST("/registrate", handlers.RegistrateHandler)
	router.POST("/refresh", handlers.RefreshHandler)
	router.POST("/auth", handlers.AuthHandler)
}

func SetupRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	configureRoutes(router)

	return router
}