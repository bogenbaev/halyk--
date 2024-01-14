package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
	"net/http"
)

func (h *Handler) initAPI2Routes(api *gin.RouterGroup) {
	crud := api.Group("/api2")
	{
		crud.GET("", h.GetAPI2Data)
		crud.PUT("", h.UpdateAPI2Data)
	}
}

func (h *Handler) GetAPI2Data(ctx *gin.Context) {
	text := ctx.Query("text")
	typeApi2 := ctx.Query("type")

	data, err := h.services.API2.Get(ctx, models.ExternalDateFact{
		Text: text,
		Type: typeApi2,
	})
	if err != nil {
		logrus.Errorf("[service error] - %+v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
	return
}

func (h *Handler) UpdateAPI2Data(ctx *gin.Context) {
	var in models.ExternalDateFact
	if err := ctx.ShouldBindUri(&in); err != nil {
		logrus.Errorf("[validation error] - %+v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := h.services.API2.Update(ctx, in); err != nil {
		logrus.Errorf("[service error] - %+v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"reason": "updated!"})
}
