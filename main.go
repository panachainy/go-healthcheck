package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-healthcheck/dto"
	"net/http"
	"os"

	oauth2ns "github.com/nmrshll/oauth2-noserver"
	"golang.org/x/oauth2"
)

func main() {
	// Submit section
	accessToken, err := getLineToken()
	if err != nil {
		panic(fmt.Sprintf("Can't get line token: ", err))
	}

	summary := &dto.Summary{
		TotalWebsites: 0,
		Success:       0,
		Failure:       0,
		TotalTime:     0,
	}

	submitReport(accessToken, summary)
}

func getHealthFromFile(csvPath string) error {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return err
	}
	for i, line := range csvLines {
		// Check header
		if i == 0 {
			if strings.ToLower(line[0]) != "url" {
				return fmt.Errorf("Invalid csv file")
			}
			continue
		}

		health := dto.Health{
			URL: line[0],
		}
		fmt.Printf("========: %v\n", health.URL)
	}

	return nil
}

func getLineToken() (string, error) {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"profile openid"},
		RedirectURL:  "http://127.0.0.1:14565/oauth/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://access.line.me/oauth2/v2.1/authorize",
			TokenURL: "https://api.line.me/oauth2/v2.1/token",
		},
	}

	client, err := oauth2ns.AuthenticateUser(conf)
	if err != nil {
		return "", err
	}

	return client.Token.AccessToken, nil
}

func submitReport(accessToken string, summary *dto.Summary) error {
	fmt.Printf("accessToken: %s\n", accessToken)

	summaryJSON, err := json.Marshal(summary)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST",
		"https://backend-challenge.line-apps.com/healthcheck/report",
		bytes.NewBuffer(summaryJSON))
	req.Header.Add(constants.CONTENT_TYPE, constants.JSON_APPLICATION)
	req.Header.Add(constants.AUTHORIZATION, "Bearer "+accessToken)

	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)

	return nil
}
