package calculate_percent_division

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"gitlab.com/a5805/ondeu/ondeu-back/external"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/repository"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
	"math"
)

type service struct {
	cfg  *models.AppConfigs
	info *external.ExternalServices
	repo *repository.Repository
}

func NewPercentDivisionService(cfg *models.AppConfigs, info *external.ExternalServices, repo *repository.Repository) *service {
	return &service{
		info: info,
		cfg:  cfg,
		repo: repo,
	}
}

// Генерация случайной строки заданной длины
func (s *service) generateRandomString(length int) (string, error) {
	// Вычисление размера буфера для случайных байт
	bufferSize := (length * 3) / 4

	// Генерация случайных байт
	randomBytes := make([]byte, bufferSize)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Кодирование в base64 и обрезка до нужной длины
	randomString := base64.URLEncoding.EncodeToString(randomBytes)[:length-1]

	return randomString, nil
}

func (s *service) CalculatePercentDivision(ctx context.Context) (*models.PercentageDivision, error) {
	s.cfg.Ratio.Total++

	api1Percent := float64(s.cfg.Ratio.Total) * s.cfg.Ratio.Api1Percent
	if math.Ceil(api1Percent) > float64(s.cfg.Ratio.Api1Count) {
		sName, err := s.generateRandomString(5)
		if err != nil {
			return s.cfg.Ratio, err
		}

		fName, err := s.generateRandomString(10)
		if err != nil {
			return s.cfg.Ratio, err
		}

		// Calling method of API1 service...
		percentage, err := s.info.LovePercentageService.Get(ctx, models.LovePercentage{
			FName: fName,
			SName: sName,
		})
		if err != nil {
			return s.cfg.Ratio, err
		}
		s.cfg.Ratio.Api1Count++

		return s.cfg.Ratio, s.repo.API1.Create(ctx, percentage)
	} else {
		// Calling method of API2 service...
		number, err := s.info.NumbersService.Get(ctx, models.DateFact{
			Fragment: true,
			Json:     true,
		})
		if err != nil {
			return s.cfg.Ratio, err
		}

		s.cfg.Ratio.Api2Count++
		return s.cfg.Ratio, s.repo.API2.Create(ctx, number)
	}
}
