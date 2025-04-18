package main

import (
	"log"

	"github.com/BikasGyawali/github-issue-cli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	cmd.Execute()
}
