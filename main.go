package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"go-healthcheck/constants"
	"go-healthcheck/dto"
	"go-healthcheck/externals"
	"net/http"
	"os"
	"strings"

	oauth2ns "github.com/nmrshll/oauth2-noserver"
	"golang.org/x/oauth2"
)

func main() {
	// // Receive csvPath from argument.
	// if len(os.Args) < 2 {
	// 	panic("Require csv path at first argument\nUsage: go-healthcheck test.csv")
	// }

	// csvPath := os.Args[1]

	// Mock instead real argument.
	csvPath := "test.csv"

	// Read csv file.
	healths, err := getHealthFromFile(csvPath)
	if err != nil {
		panic(fmt.Sprintf("Error getHealthFromFile: %v", err))
	}

	summary, err2 := externals.GetHealthSummary(healths)
	if err != nil {
		panic(fmt.Sprintf("Error GetHealthSummary: %v", err2))
	}

	fmt.Printf("summary: %#v\n", summary)

	// TODO: call http for health check follow csv file data
	fmt.Println("Perform website checking...")

	fmt.Println("Done!")

	// // Submit section
	// accessToken, err := getLineToken()
	// if err != nil {
	// 	panic(fmt.Sprintf("Can't get line token: ", err))
	// }

	// submitReport(accessToken, summary)

	fmt.Printf("Checked webistes: %v\n", summary.TotalWebsites)
	fmt.Printf("Successful websites: %v\n", summary.Success)
	fmt.Printf("Failure websites: %v\n", summary.TotalWebsites)
	fmt.Printf("Total times to finished checking website:Total times to finished checking website: %v\n", summary.TotalWebsites)
}

func getHealthFromFile(csvPath string) ([]dto.Health, error) {
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
