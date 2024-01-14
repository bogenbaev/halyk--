package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initPercentDivisionRoutes(v1)
		h.initAPI1Routes(v1)
		h.initAPI2Routes(v1)
	}
}
