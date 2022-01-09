package externals

import (
	"bytes"
	"encoding/json"
	"go-healthcheck/constants"
	"go-healthcheck/healthz/dto"
	"net/http"
	"os"

	oauth2ns "github.com/nmrshll/oauth2-noserver"
	"golang.org/x/oauth2"
)

func GetLineToken() (string, error) {
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

func SubmitReport(accessToken string, summary *dto.Summary, submitUrl string) (*http.Response, error) {
	summaryJSON, err := json.Marshal(summary)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST",
		submitUrl,
		bytes.NewBuffer(summaryJSON))
	req.Header.Add(constants.CONTENT_TYPE, constants.JSON_APPLICATION)
	req.Header.Add(constants.AUTHORIZATION, "Bearer "+accessToken)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
