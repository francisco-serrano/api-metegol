package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
)

type ParticipantController struct {
	ServiceFactory func() *ParticipantService
}

func (p *ParticipantController) AddParticipant(ctx *gin.Context) {
	rawRequest, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}

	var request AddParticipantRequest

	if err := json.Unmarshal(rawRequest, &request); err != nil {
		panic(err)
	}

	p.ServiceFactory().AddParticipant(request)

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "participant created",
	})
}

func (p *ParticipantController) GetParticipants(ctx *gin.Context) {
	participants := p.ServiceFactory().GetParticipants()

	ctx.JSON(http.StatusOK, participants)
}

// TODO: Split into /views/participant.go
type AddParticipantRequest struct {
	Name string `json:"name"`
}

type GetParticipantsResponse struct {
	Participants []string `json:"participants"`
}

// TODO: Split into /services/participant.go
type ParticipantService struct {
	Db           *gorm.DB
	Participants []string
}

func NewParticipantService(db *gorm.DB) *ParticipantService {
	return &ParticipantService{
		Db:           db,
		Participants: []string{},
	}
}

func (p *ParticipantService) AddParticipant(request AddParticipantRequest) {
	p.Participants = append(p.Participants, request.Name)
}

func (p *ParticipantService) GetParticipants() GetParticipantsResponse {
	return GetParticipantsResponse{Participants: p.Participants}
}
