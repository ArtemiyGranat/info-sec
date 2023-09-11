package routes

import (
	"info-sec-api/internal/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func configureRoutes(router *gin.Engine, db *mongo.Database, address string) {
	router.POST("/registrate", handlers.RegistrateHandler(db))
	router.POST("/refresh", handlers.RefreshHandler(db, address))
	router.POST("/auth", handlers.AuthHandler(db, address))
	router.GET("/hello", handlers.HelloHandler)
}

func SetupRouter(db *mongo.Database, address string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.SetTrustedProxies(nil)
	configureRoutes(router, db, address)

	return router
}