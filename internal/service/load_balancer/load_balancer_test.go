package load_balancer

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
)

func Test_getRandomURLByWeight_MultipleRequestWithDeviation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		totalCalls int
		balancer   []models.Balance
		deviation  float64
		wantFail   bool
	}{
		{
			name:       "1000 request with 10% deviation",
			totalCalls: 1000,
			balancer: []models.Balance{
				{
					"https://love-calculator.p.rapidapi.com/getPercentage?sname=Bob&fname=Alisa",
					80,
				},
				{
					"https://love-calculator.p.rapidapi.com/getPercentage?sname=Bob1&fname=Alisa1",
					20,
				},
			},
			deviation: 0.1,
			wantFail:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			counter := make(map[string]int)
			s := NewLoadBalancerService(&models.AppConfigs{
				Port: "",
				Cache: &models.Redis{
					"",
					"",
					"",
				},
				Balances: tt.balancer,
				LogLevel: "testing",
			}, nil, nil)

			// Act
			for i := 0; i < 1000; i++ {
				url := s.getRandomURLByWeight()
				counter[url]++
			}

			totalWeight := 0
			for _, balance := range tt.balancer {
				totalWeight += balance.Weight
			}

			for _, balance := range tt.balancer {
				calls := counter[balance.Url]
				weight := float64(balance.Weight)
				total := float64(totalWeight)
				floatTotalCalls := float64(tt.totalCalls)
				expectedLowerBoundCall := (weight / total) * floatTotalCalls * (1 - tt.deviation)
				expectedUpperBoundCall := (weight / total) * floatTotalCalls * (1 + tt.deviation)

				// Assert
				if tt.wantFail {
					require.True(t, float64(calls) > expectedUpperBoundCall)
					require.True(t, float64(calls) < expectedLowerBoundCall)
				} else {
					require.True(t, expectedLowerBoundCall < float64(calls))
					require.True(t, expectedUpperBoundCall >= float64(calls))
				}
			}
		})
	}
}
