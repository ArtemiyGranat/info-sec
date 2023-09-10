package routes

import "github.com/gin-gonic/gin"

func configureRoutes(router *gin.Engine) {
	
}

func SetupRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	configureRoutes(router)

	return router
}