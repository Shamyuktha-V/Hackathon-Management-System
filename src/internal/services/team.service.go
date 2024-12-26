package services

import (
	"Hackathon-Management-System/src/graph/model"
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/repositories"
	"context"

	"github.com/google/uuid"
)

type TeamService struct {
	AppConfig      *appConfig.AppConfig
	TeamRepository repositories.TeamRepository
}

func NewTeamService(appConfig *appConfig.AppConfig) *TeamService {
	return &TeamService{
		AppConfig:      appConfig,
		TeamRepository: *repositories.NewTeamRepository(appConfig),
	}
}

type ITeamService interface {
	GetTeam(ctx context.Context, id uuid.UUID) (*model.Team, error)
	GetTeamById(ctx context.Context, id uuid.UUID) (*model.Team, error)
	GetTeamByLeader(ctx context.Context, leaderID uuid.UUID) (*model.Team, error)
	CreateTeam(ctx context.Context, input model.CreateTeamInput) (*model.Team, error)
	UpdateTeam(ctx context.Context, id uuid.UUID, input model.UpdateTeamInput) (*model.Team, error)
	DeleteTeam(ctx context.Context, id uuid.UUID) (string, error)
}

func (s TeamService) GetTeam(ctx context.Context, id uuid.UUID) (*model.Team, error) {
	team, err := s.TeamRepository.GetTeam(ctx, id)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (s TeamService) GetTeamByLeader(ctx context.Context, leaderID uuid.UUID) (*model.Team, error) {
	team, err := s.TeamRepository.GetTeamByLeader(ctx, leaderID)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (s TeamService) CreateTeam(ctx context.Context, input model.CreateTeamInput) (*model.Team, error) {
	team, err := s.TeamRepository.CreateTeam(ctx, input)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (s TeamService) UpdateTeam(ctx context.Context, id uuid.UUID, input model.UpdateTeamInput) (*model.Team, error) {
	team, err := s.TeamRepository.UpdateTeam(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (s TeamService) DeleteTeam(ctx context.Context, id uuid.UUID) (string, error) {
	message, err := s.TeamRepository.DeleteTeam(ctx, id)
	if err != nil {
		return "", err
	}
	return message, nil
}
