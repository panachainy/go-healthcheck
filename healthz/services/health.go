package services

import (
	"encoding/csv"
	"fmt"
	"go-healthcheck/healthz/dto"
	"go-healthcheck/healthz/externals"
	"math/big"
	"os"
	"strings"
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

func GetHealthFromFile(csvPath string) ([]dto.Health, error) {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	var healthData []dto.Health
	for i, line := range csvLines {
		// Check header
		if i == 0 {
			if strings.ToLower(line[0]) != "url" {
				return nil, fmt.Errorf("Invalid csv file")
			}
			continue
		}

		healthData = append(healthData, dto.Health{URL: line[0]})
	}

	return healthData, nil
}
