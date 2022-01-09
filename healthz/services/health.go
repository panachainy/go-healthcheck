package services

import (
	"encoding/csv"
	"fmt"
	"go-healthcheck/healthz/dto"
	"go-healthcheck/healthz/externals"
	"os"
	"strings"
	"time"
)

func GetHealthSummary(healths []dto.Health) *dto.Summary {
	summary := &dto.Summary{}

	// Start time
	start := time.Now()

	for _, health := range healths {
		err := externals.Client.GetHealthCheck(health.URL)
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

	return summary
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
			if err := checkCSVHeader(line); err != nil {
				return nil, err
			}

			continue
		}

		healthData = append(healthData, dto.Health{URL: line[0]})
	}

	return healthData, nil
}

func checkCSVHeader(line []string) error {
	if strings.ToLower(line[0]) != "url" {
		return fmt.Errorf("invalid csv file: %v", "csv must have header url in first of column")
	}
	return nil
}
