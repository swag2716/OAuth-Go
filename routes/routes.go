package routes

import (
	"OAuth-Go/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	router.Handle("/", http.FileServer(http.Dir("templates")))

	// OauthGoogle
	router.HandleFunc("/auth/google/login", controllers.OauthGoogleLogin)
	router.HandleFunc("/auth/google/callback", controllers.OauthGoogleCallback)
}
