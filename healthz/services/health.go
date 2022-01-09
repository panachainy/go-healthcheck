package services

import (
	"fmt"
	"go-healthcheck/healthz/dto"
	"go-healthcheck/healthz/externals"
	"math/big"
	"time"
)

func GetHealthSummary(healths []dto.Health) (*dto.Summary, error) {
	summary := &dto.Summary{}

	// Start time
	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	for _, health := range healths {
		err := externals.GetHealthCheck(health.URL)
		if err != nil {
			summary.Failure++
			continue
		}

		summary.Success++
	}

	// End time
	elapsed := time.Since(start)

	summary.TotalWebsites = len(healths)
	summary.TotalTime = elapsed.Nanoseconds()

	return summary, nil
}
