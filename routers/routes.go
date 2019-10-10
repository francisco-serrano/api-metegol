package routers

import (
	"github.com/api-metegol/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	healthController := controllers.NewHealthController()

	router := gin.Default()
	router.GET("/health", healthController.Health)

	return router
}
