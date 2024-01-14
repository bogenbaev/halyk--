package client

import (
	"context"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
)

type ExternalServices struct {
	Proxifier
}

type Proxifier interface {
	// SendRequest gets a some information by the given address...
	SendRequest(context.Context, string) (models.Response, error)
}

func NewExternalService(cfg *models.AppConfigs) *ExternalServices {
	return &ExternalServices{
		Proxifier: NewClientService(cfg),
	}
}
