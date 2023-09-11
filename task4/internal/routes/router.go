package routes

import (
	"info-sec-api/internal/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func configureRoutes(router *gin.Engine, db *mongo.Database) {
	router.POST("/registrate", handlers.RegistrateHandler(db))
	router.POST("/refresh", handlers.RefreshHandler(db))
	router.POST("/auth", handlers.AuthHandler(db))
}

func SetupRouter(db *mongo.Database) *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	configureRoutes(router, db)

	return router
}