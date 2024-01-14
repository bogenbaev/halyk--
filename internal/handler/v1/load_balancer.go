package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
	"net/http"
)

func (h *Handler) initLoadBalancer(api *gin.RouterGroup) {
	crud := api.Group("/")
	{
		crud.Any("", h.SendRequest)
	}
}

// Custom struct to represent any type
type AnyData struct {
	Data interface{} `json:"data"`
}

func (h *Handler) SendRequest(ctx *gin.Context) {
	input := models.Request{
		IP:     ctx.ClientIP(),
		Url:    ctx.Request.URL.String(),
		Method: ctx.Request.Method,
	}

	var requestData AnyData
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Body = requestData

	response, err := h.services.Proxifier.Proxify(ctx, input)
	if err != nil {
		logrus.Errorf("[service error] - %+v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
	return
}
