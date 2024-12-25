package services

import (
	"Hackathon-Management-System/src/graph/model"
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/repositories"
	"context"

	"github.com/google/uuid"
)

type SubmissionService struct {
	AppConfig          appConfig.AppConfig
	SubmissionRepository repositories.SubmissionRepository
}

func NewSubmissionService() *SubmissionService {
	return &SubmissionService{
		AppConfig:          *appConfig.NewConfig(),
		SubmissionRepository: *repositories.NewSubmissionRepository(*appConfig.NewConfig()),
	}
}

type ISubmissionService interface {
	GetSubmission(ctx context.Context, id uuid.UUID) (*model.Submission, error)
	GetSubmissionByTeamAndHackathon(ctx context.Context, teamID uuid.UUID, hackathonID uuid.UUID) (*model.Submission, error)
	CreateSubmission(ctx context.Context, input model.CreateSubmissionInput) (*model.Submission, error)
	UpdateSubmission(ctx context.Context, id uuid.UUID, input model.UpdateSubmissionInput) (*model.Submission, error)
	DeleteSubmission(ctx context.Context, id uuid.UUID) (string, error)
}

func (s SubmissionService) GetSubmission(ctx context.Context, id uuid.UUID) (*model.Submission, error) {
	submission, err := s.SubmissionRepository.GetSubmission(ctx, id)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (s SubmissionService) GetSubmissionByTeamAndHackathon(ctx context.Context, teamID uuid.UUID, hackathonID uuid.UUID) (*model.Submission, error) {
	submission, err := s.SubmissionRepository.GetSubmissionByTeamAndHackathon(ctx, teamID, hackathonID)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (s SubmissionService) CreateSubmission(ctx context.Context, input model.CreateSubmissionInput) (*model.Submission, error) {
	submission, err := s.SubmissionRepository.CreateSubmission(ctx, input)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (s SubmissionService) UpdateSubmission(ctx context.Context, id uuid.UUID, input model.UpdateSubmissionInput) (*model.Submission, error) {
	submission, err := s.SubmissionRepository.UpdateSubmission(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (s SubmissionService) DeleteSubmission(ctx context.Context, id uuid.UUID) (string, error) {
	message, err := s.SubmissionRepository.DeleteSubmission(ctx, id)
	if err != nil {
		return "", err
	}
	return message, nil
}
