package services

import (
	"Hackathon-Management-System/src/graph/model"
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/repositories"
	"context"

	"github.com/google/uuid"
)

type HackathonService struct {
	AppConfig           *appConfig.AppConfig
	HackathonRepository repositories.HackathonRepository
}

func NewHackathonService(appConfig *appConfig.AppConfig) *HackathonService {
	return &HackathonService{
		AppConfig:           appConfig,
		HackathonRepository: *repositories.NewHackathonRepository(appConfig),
	}
}

type IHackathonService interface {
	GetHackathon(ctx context.Context, id uuid.UUID) (*model.Hackathon, error)
	GetHackathonByCategory(ctx context.Context, categoryID uuid.UUID) (*model.Hackathon, error)
	GetHackathonsByAttributes(ctx context.Context, judgeID *uuid.UUID, name *string, problemStatement *string, startDate *string, endDate *string, categoryID *uuid.UUID, duaration *int) ([]*model.Hackathon, error)
	CreateHackathon(ctx context.Context, input model.CreateHackathonInput) (*model.Hackathon, error)
	UpdateHackathon(ctx context.Context, id uuid.UUID, input model.UpdateHackathonInput) (*model.Hackathon, error)
	DeleteHackathon(ctx context.Context, id uuid.UUID) (string, error)
	GetHackathonsAvailableForRegistration(ctx context.Context) ([]*model.Hackathon, error)
	GetFutureHackathons(ctx context.Context) ([]*model.Hackathon, error)
}

func (s HackathonService) GetHackathon(ctx context.Context, id uuid.UUID) (*model.Hackathon, error) {
	hackathon, err := s.HackathonRepository.GetHackathon(ctx, id)
	if err != nil {
		return nil, err
	}
	return hackathon, nil
}

func (s HackathonService) GetHackathonByCategory(ctx context.Context, categoryID uuid.UUID) (*model.Hackathon, error) {
	hackathon, err := s.HackathonRepository.GetHackathonByCategory(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	return hackathon, nil
}

func (s HackathonService) GetHackathonsByAttributes(ctx context.Context, judgeID *uuid.UUID, name *string, problemStatement *string, startDate *string, endDate *string, categoryID *uuid.UUID, duration *int) ([]*model.Hackathon, error) {
	hackathons, err := s.HackathonRepository.GetHackathonsByAttributes(ctx, judgeID, name, problemStatement, startDate, endDate, categoryID, duration)
	if err != nil {
		return nil, err
	}
	return hackathons, nil
}

func (s HackathonService) CreateHackathon(ctx context.Context, input model.CreateHackathonInput) (*model.Hackathon, error) {
	hackathon, err := s.HackathonRepository.CreateHackathon(ctx, input)
	if err != nil {
		return nil, err
	}
	return hackathon, nil
}

func (s HackathonService) UpdateHackathon(ctx context.Context, id uuid.UUID, input model.UpdateHackathonInput) (*model.Hackathon, error) {
	hackathon, err := s.HackathonRepository.UpdateHackathon(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return hackathon, nil
}

func (s HackathonService) DeleteHackathon(ctx context.Context, id uuid.UUID) (string, error) {
	message, err := s.HackathonRepository.DeleteHackathon(ctx, id)
	if err != nil {
		return "", err
	}
	return message, nil
}

func (s HackathonService) GetHackathonsAvailableForRegistration(ctx context.Context) ([]*model.Hackathon, error) {
	hackathons, err := s.HackathonRepository.GetHackathonsAvailableForRegistration(ctx)
	if err != nil {
		return nil, err
	}
	return hackathons, nil
}

func (s HackathonService) GetFutureHackathons(ctx context.Context) ([]*model.Hackathon, error) {
	hackathons, err := s.HackathonRepository.GetFutureHackathons(ctx)
	if err != nil {
		return nil, err
	}
	return hackathons, nil
}
