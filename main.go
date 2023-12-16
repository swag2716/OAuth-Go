package main

import (
	"OAuth-Go/helpers"
	"OAuth-Go/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	helpers.LoadEnv()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("port error : ", port)
		// port = "8000"
	}

	router := mux.NewRouter()

	routes.Routes(router)

	fmt.Println("Server is running on port :", port)

	http.ListenAndServe(":"+port, router)

}
