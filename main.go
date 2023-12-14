package main

import (
	"OAuth-Go/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatal("error: ", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal(port)
	}

	router := mux.NewRouter()

	routes.Routes(router)

	http.ListenAndServe(":"+port, router)

}
