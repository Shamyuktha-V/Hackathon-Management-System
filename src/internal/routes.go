package internal

import (
	"Hackathon-Management-System/src/rest/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	googleCallbackUrl := GoogleCallbackURL
	googleController := controllers.NewGoogleController()

	router.GET(googleCallbackUrl, func(c *gin.Context) {
		googleController.ProcessOAuthForGoogle(c)
	})

	return router
}
