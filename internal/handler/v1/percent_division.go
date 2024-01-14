package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/repository"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models/dto"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/utils"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) initPercentDivisionRoutes(api *gin.RouterGroup) {
	crud := api.Group("/calc")
	{
		crud.GET("", h.)
	}
}

func (h *Handler) GetPercentDivision(ctx *gin.Context) {
	 percentDivision, err := h.services.CalculatePercentDivision.CalculatePercentDivision(ctx)
	 if err != nil {
		 logrus.Errorf("[validaton error] - %+v", err)
		 ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
}
