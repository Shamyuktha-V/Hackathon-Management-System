package controllers

import (
	"Hackathon-Management-System/src/rest/services"

	"github.com/gin-gonic/gin"
)

type GoogleController struct {
	GoogleService *services.GoogleService
}

func NewGoogleController() *GoogleController {
	return &GoogleController{
		GoogleService: services.NewGoogleServices(),
	}
}

func (c GoogleController) ProcessOAuthForGoogle(ctx *gin.Context) {
	user, err := c.GoogleService.ProcessOAuth(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"user": user,
		})
	}
}
