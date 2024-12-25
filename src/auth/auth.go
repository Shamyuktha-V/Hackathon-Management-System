package auth

import (
	"Hackathon-Management-System/src/auth/utils"
	"Hackathon-Management-System/src/graph/model"
	"Hackathon-Management-System/src/internal/config"
	services "Hackathon-Management-System/src/internal/services"
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GetUser(ctx context.Context, appConfig *config.AppConfig) (*model.User, error) {
	jwtToken := ctx.Value("Authorization").(string)
	if strings.HasPrefix(jwtToken, "Bearer ") {
		jwtToken = strings.TrimPrefix(jwtToken, "Bearer ")
	}
	print("Jwt token :: ", jwtToken)
	userID, err := utils.VerifyJWTToken(jwtToken)
	if err != nil {
		return nil, err
	}

	userService := services.NewUserService(appConfig)
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	fmt.Print("User UUID :: ", userUUID)
	user, err := userService.GetUser(ctx, userUUID)
	if err != nil {
		return nil, err
	}
	fmt.Print("User :: ", user)
	return user, nil
}
