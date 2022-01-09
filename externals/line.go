package externals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-healthcheck/constants"
	"go-healthcheck/dto"
	"net/http"
	"os"

	oauth2ns "github.com/nmrshll/oauth2-noserver"
	"golang.org/x/oauth2"
)

func SubmitReport(summary *dto.Summary) {
	// Submit section
	accessToken, err := getLineToken()
	if err != nil {
		panic(fmt.Sprintf("Can't get line token: %v", err))
	}

	submitReport(accessToken, summary)
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
