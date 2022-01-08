package externals

import (
	"fmt"
	"go-healthcheck/dto"
	"net/http"
	"time"
)

func GetHealthSummary(healths []dto.Health) (*dto.Summary, error) {
	summary := &dto.Summary{}

	for _, health := range healths {
		err := getHealthCheck(health.URL)
		fmt.Printf("URL: %s Error: %v\n", health.URL, err)
		if err != nil {
			summary.Failure++
			continue
		}

		summary.Success++
	}

	summary.TotalWebsites = len(healths)

	// TODO: TotalTime

	return summary, nil
}

func getHealthCheck(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 1 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()

	return nil
}
