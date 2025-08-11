// Package initializer loads the initial packages for the app
package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

// EnvVariables loads the environment variables for the app
func EnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Unable to load env variables: %v\n", err.Error())
	}
}
