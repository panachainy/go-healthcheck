package main

import (
	"log"
	"os"

	oauth2ns "github.com/nmrshll/oauth2-noserver"
	"golang.org/x/oauth2"
)

func main() {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"profile openid"},
		RedirectURL:  "http://localhost:14565/oauth/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://access.line.me/oauth2/v2.1/authorize",
			TokenURL: "https://access.line.me/oauth2/v2.1/access_token",
		},
	}

	client, err := oauth2ns.AuthenticateUser(conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Before authenticated")
	// use client.Get / client.Post for further requests, the token will automatically be there
	_, _ = client.Get("/auth-protected-path")
	log.Printf("Successfully authenticated")
}
