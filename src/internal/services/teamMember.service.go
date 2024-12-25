package services

import (
	"Hackathon-Management-System/src/graph/model"
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/repositories"
	"context"

	"github.com/google/uuid"
)

type TeamMemberService struct {
	AppConfig          appConfig.AppConfig
	TeamMemberRepository repositories.TeamMemberRepository
}

func NewTeamMemberService() *TeamMemberService {
	return &TeamMemberService{
		AppConfig:          *appConfig.NewConfig(),
		TeamMemberRepository: *repositories.NewTeamMemberRepository(*appConfig.NewConfig()),
	}
}

type ITeamMemberService interface {
	GetTeamMember(ctx context.Context, id uuid.UUID) (*model.TeamMember, error)
	GetTeamMembersByTeamID(ctx context.Context, teamID uuid.UUID) ([]*model.TeamMember, error)
	GetTeamMembersByUserID(ctx context.Context, userID uuid.UUID) ([]*model.TeamMember, error)
	CreateTeamMember(ctx context.Context, input model.CreateTeamMemberInput) (*model.TeamMember, error)
	UpdateTeamMember(ctx context.Context, id uuid.UUID, input model.UpdateTeamMemberInput) (*model.TeamMember, error)
	DeleteTeamMember(ctx context.Context, id uuid.UUID) (string, error)
}

func (s TeamMemberService) GetTeamMember(ctx context.Context, id uuid.UUID) (*model.TeamMember, error) {
	teamMember, err := s.TeamMemberRepository.GetTeamMember(ctx, id)
	if err != nil {
		return nil, err
	}
	return teamMember, nil
}

func (s TeamMemberService) GetTeamMembersByTeamID(ctx context.Context, teamID uuid.UUID) ([]*model.TeamMember, error) {
	teamMembers, err := s.TeamMemberRepository.GetTeamMembersByTeamID(ctx, teamID)
	if err != nil {
		return nil, err
	}
	return teamMembers, nil
}

func (s TeamMemberService) GetTeamMembersByUserID(ctx context.Context, userID uuid.UUID) ([]*model.TeamMember, error) {
	teamMembers, err := s.TeamMemberRepository.GetTeamMembersByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return teamMembers, nil
}

func (s TeamMemberService) CreateTeamMember(ctx context.Context, input model.CreateTeamMemberInput) (*model.TeamMember, error) {
	teamMember, err := s.TeamMemberRepository.CreateTeamMember(ctx, input)
	if err != nil {
		return nil, err
	}
	return teamMember, nil
}

func (s TeamMemberService) UpdateTeamMember(ctx context.Context, id uuid.UUID, input model.UpdateTeamMemberInput) (*model.TeamMember, error) {
	teamMember, err := s.TeamMemberRepository.UpdateTeamMember(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return teamMember, nil
}

func (s TeamMemberService) DeleteTeamMember(ctx context.Context, id uuid.UUID) (string, error) {
	message, err := s.TeamMemberRepository.DeleteTeamMember(ctx, id)
	if err != nil {
		return "", err
	}
	return message, nil
}
