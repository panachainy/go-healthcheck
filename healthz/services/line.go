package services

import (
	"fmt"
	"go-healthcheck/healthz/dto"
	"go-healthcheck/healthz/externals"
)

func SubmitReport(summary *dto.Summary) {
	accessToken, err := externals.GetLineToken()
	if err != nil {
		panic(fmt.Sprintf("Can't get line token: %v", err))
	}

	_, err = externals.SubmitReport(accessToken, summary)
	if err != nil {
		panic(fmt.Sprintf("Submit Error: %v", err))
	}
}
