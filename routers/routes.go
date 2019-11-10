package routers

import (
	"github.com/api-metegol/controllers"
	"github.com/api-metegol/utils"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, deps utils.Dependencies) {
	healthController := controllers.NewHealthController()

	participantsController := controllers.ParticipantController{
		ServiceFactory: func() *controllers.ParticipantService {
			return controllers.NewParticipantService(deps.Db)
		},
	}

	tournamentsController := controllers.TournamentController{
		ServiceFactory: func() *controllers.TournamentService {
			return controllers.NewTournamentService(deps.Db)
		},
	}

	// Health Check
	router.GET("/health", healthController.Health)

	// Participants
	router.POST("/participants", participantsController.AddParticipant)
	router.GET("/participants", participantsController.GetParticipants)

	// Tournaments
	router.POST("/tournaments", tournamentsController.AddTournament)
	router.GET("/tournaments", tournamentsController.GetTournaments)
}
