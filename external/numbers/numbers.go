package numbers

import (
	"context"
	"encoding/json"
	"fmt"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
	"io"
	"net/http"
)

type service struct {
	cfg *models.AppConfigs
}

func NewNumberService(cfg *models.AppConfigs) *service {
	return &service{
		cfg: cfg,
	}
}

func (s *service) Get(_ context.Context, input models.DateFact) (out models.ExternalDateFact, err error) {
	url := s.cfg.API2 + fmt.Sprintf("/6/21/date?fragment=%v&json=%v", input.Fragment, input.Json)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "c82637c962mshb6672571c2f05c8p1b610fjsn7eb0cf6db9e8")
	req.Header.Add("X-RapidAPI-Host", "numbersapi.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	if err = json.Unmarshal(body, &out); err != nil {
		return out, err
	}

	return out, nil
}
