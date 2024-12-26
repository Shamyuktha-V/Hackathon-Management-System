package graph

import (
	appConfig "Hackathon-Management-System/src/internal/config"
	service "Hackathon-Management-System/src/internal/services"
)

type Resolver struct {
	AppConfig         *appConfig.AppConfig
	QueryResolver     queryResolver
	MutationResolver  mutationResolver
	UserService       *service.UserService
	SubmissionService *service.SubmissionService
	HackathonService  *service.HackathonService
	CategoryService   *service.CategoryService
	TeamMemberService *service.TeamMemberService
	TeamService       *service.TeamService
}

func NewResolver(appConfig *appConfig.AppConfig) *Resolver {
	return &Resolver{
		AppConfig:         appConfig,
		QueryResolver:     queryResolver{},
		MutationResolver:  mutationResolver{},
		UserService:       service.NewUserService(appConfig),
		SubmissionService: service.NewSubmissionService(appConfig),
		HackathonService:  service.NewHackathonService(appConfig),
		CategoryService:   service.NewCategoryService(appConfig),
		TeamMemberService: service.NewTeamMemberService(appConfig),
		TeamService:       service.NewTeamService(appConfig),
	}
}
