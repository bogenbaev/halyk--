package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
	"net/http"
)

func (h *Handler) initAPI1Routes(api *gin.RouterGroup) {
	crud := api.Group("/api1")
	{
		crud.GET("", h.GetAPI1Data)
		crud.PUT("", h.UpdateAPI1Data)
	}
}

func (h *Handler) GetAPI1Data(ctx *gin.Context) {
	sName := ctx.Query("sName")
	fName := ctx.Query("fName")

	data, err := h.services.API1.Get(ctx, models.ExternalLovePercentage{
		SName: sName,
		FName: fName,
	})
	if err != nil {
		logrus.Errorf("[service error] - %+v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
	return
}

func (h *Handler) UpdateAPI1Data(ctx *gin.Context) {
	var in models.ExternalLovePercentage
	if err := ctx.ShouldBindUri(&in); err != nil {
		logrus.Errorf("[validation error] - %+v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	if err := h.services.API1.Update(ctx, in); err != nil {
		logrus.Errorf("[service error] - %+v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"reason": "updated!"})
}
