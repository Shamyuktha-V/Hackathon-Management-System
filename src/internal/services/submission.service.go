package services

import (
	"Hackathon-Management-System/src/graph/model"
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/repositories"
	"context"

	"github.com/google/uuid"
)

type SubmissionService struct {
	AppConfig            *appConfig.AppConfig
	SubmissionRepository repositories.SubmissionRepository
}

func NewSubmissionService(appConfig *appConfig.AppConfig) *SubmissionService {
	return &SubmissionService{
		AppConfig:            appConfig,
		SubmissionRepository: *repositories.NewSubmissionRepository(appConfig),
	}
}

type ISubmissionService interface {
	GetSubmission(ctx context.Context, id uuid.UUID) (*model.Submission, error)
	GetSubmissionByTeamAndHackathon(ctx context.Context, teamID uuid.UUID, hackathonID uuid.UUID) (*model.Submission, error)
	CreateSubmission(ctx context.Context, input model.CreateSubmissionInput) (*model.Submission, error)
	UpdateSubmission(ctx context.Context, id uuid.UUID, input model.UpdateSubmissionInput) (*model.Submission, error)
	DeleteSubmission(ctx context.Context, id uuid.UUID) (string, error)
	GetHackathonsForAJudge(ctx context.Context, judgeID uuid.UUID) ([]*model.Submission, error)
	GetHackathonsForATeam(ctx context.Context, teamID uuid.UUID) ([]*model.Submission, error)
	GetTeamsForAHackathons(ctx context.Context, hackathonID uuid.UUID) ([]*model.Submission, error)
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

func (s SubmissionService) GetHackathonsForAJudge(ctx context.Context, judgeID uuid.UUID) ([]*model.Submission, error) {
	submissions, err := s.SubmissionRepository.GetHackathonsForAJudge(ctx, judgeID)
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func (s SubmissionService) GetHackathonsForATeam(ctx context.Context, teamID uuid.UUID) ([]*model.Submission, error) {
	submissions, err := s.SubmissionRepository.GetHackathonsForATeam(ctx, teamID)
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func (s SubmissionService) GetTeamsForAHackathons(ctx context.Context, hackathonID uuid.UUID) ([]*model.Submission, error) {
	submissions, err := s.SubmissionRepository.GetTeamsForAHackathons(ctx, hackathonID)
	if err != nil {
		return nil, err
	}
	return submissions, nil
}
