package services

import (
	"fmt"
	"go-healthcheck/healthz/dto"
	"go-healthcheck/healthz/externals"
)

func SubmitReport(summary *dto.Summary) {
	// Submit section
	accessToken, err := externals.GetLineToken()
	if err != nil {
		panic(fmt.Sprintf("Can't get line token: %v", err))
	}

	externals.SubmitReport(accessToken, summary)
}
