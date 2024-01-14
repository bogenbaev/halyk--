package external

import (
	"context"
	"gitlab.com/a5805/ondeu/ondeu-back/external/love_percentage"
	"gitlab.com/a5805/ondeu/ondeu-back/external/numbers"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
)

type ExternalServices struct {
	LovePercentageService
	NumbersService
}

type LovePercentageService interface {
	// Get returns a love percentage information
	Get(context.Context, models.LovePercentage) (models.ExternalLovePercentage, error)
}

type NumbersService interface {
	// Get returns a date fact from world history
	Get(context.Context, models.DateFact) (models.ExternalDateFact, error)
}

func NewExternalService(cfg *models.AppConfigs) *ExternalServices {
	return &ExternalServices{
		LovePercentageService: love_percentage.NewLoveService(cfg),
		NumbersService:        numbers.NewNumberService(cfg),
	}
}
