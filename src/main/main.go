package main

import (
	"fmt"

	"log"

	"Hackathon-Management-System/src/internal/app"
	configuration "Hackathon-Management-System/src/internal/config"

	"github.com/joho/godotenv"
)

var appConfig *configuration.AppConfig

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Println("error loading env file"+" : %s", err)
	}

	appConfig = configuration.NewConfig()
}

func main() {
	fmt.Println("Connecting to application")
	app.StartApplication(appConfig)
}
