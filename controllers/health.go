package controllers

import (
	"github.com/api-metegol/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, views.NewHealthResponse("metegol service"))
}
