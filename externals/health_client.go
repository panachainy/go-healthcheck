package externals

import (
	"fmt"
	"go-healthcheck/dto"
	"math/big"
	"net/http"
	"time"
)

func GetHealthSummary(healths []dto.Health) (*dto.Summary, error) {
	summary := &dto.Summary{}

	// Start time
	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	for _, health := range healths {
		err := getHealthCheck(health.URL)
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
