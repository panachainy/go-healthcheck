package services

import (
	"fmt"
	lineConstants "go-healthcheck/healthz/constants"
	"go-healthcheck/healthz/dto"
	"go-healthcheck/healthz/externals"
)

func SubmitReport(summary *dto.Summary) {
	accessToken, err := externals.GetLineToken()
	if err != nil {
		panic(fmt.Sprintf("Can't get line token: %v", err))
	}

	_, err = externals.SubmitReport(accessToken, summary, lineConstants.LINE_SUBMIT_URL)
	if err != nil {
		panic(fmt.Sprintf("Submit Error: %v", err))
	}
}
