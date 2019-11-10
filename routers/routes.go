package routers

import (
	"fmt"
	"github.com/api-metegol/controllers"
	"github.com/api-metegol/utils"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, deps utils.Dependencies) {
	fmt.Println("AAAAAAAAAAAA", deps)

	healthController := controllers.NewHealthController()
	participantsController := controllers.NewParticipantController()
	tournamentsController := controllers.NewTournamentController()

	// Health Check
	router.GET("/health", healthController.Health)

	// Participants
	router.POST("/participants", participantsController.AddParticipant)
	router.GET("/participants", participantsController.GetParticipants)

	// Tournaments
	router.POST("/tournaments", tournamentsController.AddTournament)
	router.GET("/tournaments", tournamentsController.GetTournaments)
}
