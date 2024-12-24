package main

import (
	"Hackathon-Management-System/src/internal"
	"log"
)

func main() {
	router := internal.SetupRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
