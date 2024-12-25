package graph

import (
	appConfig "Hackathon-Management-System/src/internal/config"
	service "Hackathon-Management-System/src/internal/services"
)

type Resolver struct {
	AppConfig        *appConfig.AppConfig
	QueryResolver    queryResolver
	MutationResolver mutationResolver
	UserService      *service.UserService
}

func NewResolver(appConfig *appConfig.AppConfig) *Resolver {
	return &Resolver{
		AppConfig:        appConfig,
		QueryResolver:    queryResolver{},
		MutationResolver: mutationResolver{},
		UserService:      service.NewUserService(appConfig),
	}
}
