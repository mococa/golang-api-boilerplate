package main

import (
	"log"
	"net/http"

	api "github.com/mococa/golang-api-boilerplate/server"

	"github.com/gorilla/mux"
)

func main() {
	const PORT string = ":4321"

	router := mux.NewRouter()

	api.Init()
	api.SetupRoutes(router)

	log.Println("Server listening on port", PORT)
	log.Fatalln(http.ListenAndServe(PORT, router))
}
