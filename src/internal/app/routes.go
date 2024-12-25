package app

import (
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/rest/controllers"

	"github.com/gin-gonic/gin"
)

func mapURLs(appConfig *appConfig.AppConfig) {
	googleCallbackUrl := GoogleCallbackURL
	googleController := controllers.NewGoogleController(appConfig)

	router.GET(googleCallbackUrl, func(c *gin.Context) {
		googleController.ProcessOAuthForGoogle(c)
	})
}
