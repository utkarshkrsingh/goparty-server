package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/utkarshkrsingh/goparty/internal/initializer"
	routehandlers "github.com/utkarshkrsingh/goparty/internal/routeHandlers"
)

func init() {
	initializer.EnvVariables()
}

func main() {
	router := mux.NewRouter()
	routehandlers.HandleRoutes(router)

	log.Printf("Server is starting on :%v", os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("APP_PORT"), router); err != nil {
		log.Fatalf("Server unable to start: %v\n", err)
	}
}
