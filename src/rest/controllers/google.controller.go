package controllers

import (
	"Hackathon-Management-System/src/rest/services"

	appConfig "Hackathon-Management-System/src/internal/config"

	"github.com/gin-gonic/gin"
)

type GoogleController struct {
	AppConfig     *appConfig.AppConfig
	GoogleService *services.GoogleService
}

func NewGoogleController(appConfig *appConfig.AppConfig) *GoogleController {
	return &GoogleController{
		AppConfig:     appConfig,
		GoogleService: services.NewGoogleServices(appConfig),
	}
}

func (c GoogleController) ProcessOAuthForGoogle(ctx *gin.Context) {
	user, newSignUp, err := c.GoogleService.ProcessOAuth(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"user":      user,
			"newSignUp": newSignUp,
		})
	}
}
