package configs

import (
	"OAuth-Go/helpers"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
const OauthGithubUrlAPI = "https://api.github.com/user"

var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8000/auth/google/callback",
	ClientID:     helpers.GetGoogleClientID(),
	ClientSecret: helpers.GetGoogleClientSecret(),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

var GithubOauthConfig = &oauth2.Config{
	ClientID:     helpers.GetGithubClientID(),
	ClientSecret: helpers.GetGithubClientSecret(),
	RedirectURL:  "http://localhost:8000/auth/github/callback",
	Scopes:       []string{"user:email"},
	Endpoint:     github.Endpoint,
}
