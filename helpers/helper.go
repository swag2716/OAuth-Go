package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func GenerateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error: ", err)
	}
}

func GetGoogleClientID() string {

	LoadEnv()

	githubClientID, exists := os.LookupEnv("GOOGLE_OAUTH_CLIENT_ID")
	if !exists {
		log.Fatal("Google Client ID not defined in .env file")
	}

	return githubClientID
}

func GetGoogleClientSecret() string {

	githubClientSecret, exists := os.LookupEnv("GOOGLE_OAUTH_CLIENT_SECRET")
	if !exists {
		log.Fatal("Google Client secret not defined in .env file")
	}

	return githubClientSecret
}

func GetGithubClientID() string {

	githubClientID, exists := os.LookupEnv("GITHUB_OAUTH_CLIENT_ID")
	if !exists {
		log.Fatal("Github Client ID not defined in .env file")
	}

	return githubClientID
}

func GetGithubClientSecret() string {

	githubClientSecret, exists := os.LookupEnv("GITHUB_OAUTH_CLIENT_SECRET")
	if !exists {
		log.Fatal("Github Client secret not defined in .env file")
	}

	return githubClientSecret
}
