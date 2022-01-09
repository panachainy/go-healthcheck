package externals

import (
	"go-healthcheck/constants"
	"net/http"
)

var Client IHealthService = healthService{}

type IHealthService interface {
	GetHealthCheck(url string) error
}

type healthService struct{}

func (hs healthService) GetHealthCheck(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: constants.CLIENT_TIME_OUT}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()

	return nil
}
