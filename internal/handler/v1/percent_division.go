package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) initPercentDivisionRoutes(api *gin.RouterGroup) {
	crud := api.Group("/calc")
	{
		crud.GET("", h.GetPercentDivision)
	}
}

func (h *Handler) GetPercentDivision(ctx *gin.Context) {
	percentDivision, err := h.services.CalculatePercentDivision.CalculatePercentDivision(ctx)
	if err != nil {
		logrus.Errorf("[service error] - %+v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, percentDivision)
	return
}
