package controllers

import (
	"OAuth-Go/configs"
	"OAuth-Go/helpers"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

func OauthGithubLogin(w http.ResponseWriter, r *http.Request) {
	// Create oauthState cookie
	oauthState := helpers.GenerateStateOauthCookie(w)

	u := configs.GithubOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func OauthGithubCallback(w http.ResponseWriter, r *http.Request) {
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth github state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGithub(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "UserInfo: %s\n", data)
}

func getUserDataFromGithub(code string) ([]byte, error) {
	token, err := configs.GithubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	client := configs.GithubOauthConfig.Client(context.Background(), token)
	response, err := client.Get(configs.OauthGithubUrlAPI)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}
