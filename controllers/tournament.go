package controllers

import (
	"encoding/json"
	"github.com/api-metegol/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

type TournamentController struct {
	ServiceFactory func() *TournamentService
}

func (t *TournamentController) AddTournament(ctx *gin.Context) {
	rawRequest, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}

	var request AddTournamentRequest

	if err := json.Unmarshal(rawRequest, &request); err != nil {
		panic(err)
	}

	response := t.ServiceFactory().AddTournament(request)

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":    "tournament created",
		"result": response,
	})
}

func (t *TournamentController) GetTournaments(ctx *gin.Context) {
	tournaments := t.ServiceFactory().GetTournaments()

	ctx.JSON(http.StatusOK, tournaments)
}

// TODO: Split into /services/tournament.go
type TournamentService struct {
	Db *gorm.DB
}

func NewTournamentService(db *gorm.DB) *TournamentService {
	return &TournamentService{
		Db: db,
	}
}

func (t *TournamentService) AddTournament(request AddTournamentRequest) AddTournamentResponse {
	tournament := models.Tournament{
		Name: request.Name,
	}

	if err := t.Db.Create(&tournament).Error; err != nil {
		panic(err)
	}

	return AddTournamentResponse{
		ID:        tournament.ID,
		Name:      tournament.Name,
		CreatedAt: tournament.CreatedAt,
	}
}

func (t *TournamentService) GetTournaments() GetTournamentsResponse {
	var tournaments []models.Tournament

	if err := t.Db.Find(&tournaments).Error; err != nil {
		panic(err)
	}

	response := GetTournamentsResponse{
		Tournaments: []GetTournamentResponse{},
	}

	for _, t := range tournaments {
		response.Tournaments = append(response.Tournaments, GetTournamentResponse{
			ID:        t.ID,
			Name:      t.Name,
			CreatedAt: t.CreatedAt,
		})
	}

	return response
}

// TODO: Split into /views/tournament.go
type AddTournamentRequest struct {
	Name string `json:"name"`
}

type AddTournamentResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type GetTournamentResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type GetTournamentsResponse struct {
	Tournaments []GetTournamentResponse `json:"tournaments"`
}
