package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	v1 "gitlab.com/a5805/ondeu/ondeu-back/internal/handler/v1"
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

func (h *Handler) Init() *gin.Engine {
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/health"},
	}))

	router.Use(gin.Recovery())

	// third party handlers
	router.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "Up"}) })

	h.InitRoutes(router)

	return router
}

func (h *Handler) InitRoutes(router *gin.Engine) {
	handler := v1.NewHandler(h.services)

	api := router.Group("/api")
	{
		handler.Init(api)
	}
}
