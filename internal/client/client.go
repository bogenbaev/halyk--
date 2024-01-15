package client

import (
	"context"
	"encoding/json"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
	"io"
	"net/http"
)

type service struct {
	cfg *models.AppConfigs
}

func NewClientService(cfg *models.AppConfigs) *service {
	return &service{
		cfg: cfg,
	}
}

func (s *service) SendRequest(_ context.Context, url string) (out models.Response, err error) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "c82637c962mshb6672571c2f05c8p1b610fjsn7eb0cf6db9e8")
	req.Header.Add("X-RapidAPI-Host", "love-calculator.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	resp := struct {
		FName      string `json:"fname"`
		SName      string `json:"sname"`
		Percentage string `json:"percentage"`
		Result     string `json:"result"`
	}{}

	if err = json.Unmarshal(body, &resp); err != nil {
		return out, err
	}

	out.Body = resp
	out.Status = res.Status
	return out, nil
}
