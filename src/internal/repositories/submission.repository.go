package repositories

import (
	"Hackathon-Management-System/src/graph/model"
	"Hackathon-Management-System/src/internal/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubmissionRepository struct {
	DB *gorm.DB
}

func InputToEntitySubmission(input interface{}) models.Submission {
	switch input := input.(type) {
	case model.CreateSubmissionInput:
		var entity models.Submission
		entity.TeamID = input.TeamID
		entity.HackathonID = input.HackathonID
		entity.GithubLink = input.GithubLink
		entity.DocumentURL = input.DocumentURL
		entity.PresentationURL = input.PresentationURL
		entity.SubmittedAt = input.SubmittedAt
		entity.IsSubmitted = input.IsSubmitted
		entity.KeyFeatures = input.KeyFeatures
		entity.Feedback = input.Feedback
		entity.Adherence = input.Adherence
		entity.InnovationScore = input.InnovationScore
		entity.FeasibilityScore = input.FeasibilityScore
		entity.ImpactScore = input.ImpactScore
		entity.Summary = input.Summary
		return entity

	case model.UpdateSubmissionInput:
		var entity models.Submission
		if input.TeamID != nil {
			entity.TeamID = *input.TeamID
		}
		if input.HackathonID != nil {
			entity.HackathonID = *input.HackathonID
		}
		if input.GithubLink != nil {
			entity.GithubLink = *input.GithubLink
		}
		if input.DocumentURL != nil {
			entity.DocumentURL = *input.DocumentURL
		}
		if input.PresentationURL != nil {
			entity.PresentationURL = *input.PresentationURL
		}
		if input.SubmittedAt != nil {
			entity.SubmittedAt = *input.SubmittedAt
		}
		if input.IsSubmitted != nil {
			entity.IsSubmitted = *input.IsSubmitted
		}
		if input.KeyFeatures != nil {
			entity.KeyFeatures = *input.KeyFeatures
		}
		if input.Feedback != nil {
			entity.Feedback = *input.Feedback
		}
		if input.Adherence != nil {
			entity.Adherence = *input.Adherence
		}
		if input.InnovationScore != nil {
			entity.InnovationScore = *input.InnovationScore
		}
		if input.FeasibilityScore != nil {
			entity.FeasibilityScore = *input.FeasibilityScore
		}
		if input.ImpactScore != nil {
			entity.ImpactScore = *input.ImpactScore
		}
		if input.Summary != nil {
			entity.Summary = *input.Summary
		}
		return entity
	}
	return models.Submission{}
}

func NewSubmissionRepository(db *gorm.DB) *SubmissionRepository {
	return &SubmissionRepository{DB: db}
}

func (repo *SubmissionRepository) GetSubmission(ctx context.Context, id uuid.UUID) (*model.Submission, error) {
	var submission *model.Submission

	result := repo.DB.Table(models.Submission{}.TableName()).Where("submission_id =?", id).First(&submission)
	if result.Error != nil {
		return submission, result.Error
	}
	return submission, nil
}

func (repo *SubmissionRepository) GetSubmissionByTeamAndHackathon(ctx context.Context, teamID uuid.UUID, hackathonID uuid.UUID) (*model.Submission, error) {
	var submission *model.Submission

	result := repo.DB.Table(models.Submission{}.TableName()).Where("team_id =? AND hackathon_id =?", teamID, hackathonID).First(&submission)
	if result.Error != nil {
		return submission, result.Error
	}
	return submission, nil
}

func (repo *SubmissionRepository) CreateSubmission(ctx context.Context, input model.CreateSubmissionInput) (*model.Submission, error) {
	submission := InputToEntitySubmission(input)

	result := repo.DB.Table(models.Submission{}.TableName()).Create(&submission)
	if result.Error != nil {
		return nil, result.Error
	}

	return repo.GetSubmission(ctx, submission.SubmissionID)
}

func (repo *SubmissionRepository) UpdateSubmission(ctx context.Context, id uuid.UUID, input model.UpdateSubmissionInput) (*model.Submission, error) {
	submission := InputToEntitySubmission(input)
	query := fmt.Sprintf("UPDATE %s SET version = version + 1 WHERE submission_id =?", models.Submission{}.TableName())

	if err := repo.DB.Exec(query, id).Error; err != nil {
		return nil, err
	}

	result := repo.DB.Table(models.Submission{}.TableName()).Where("submission_id =?", id).Updates(submission)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.GetSubmission(ctx, id)
}

func (repo *SubmissionRepository) DeleteSubmission(ctx context.Context, id uuid.UUID) (string, error) {
	var submission *model.Submission

	result := repo.DB.Table(models.Submission{}.TableName()).Where("submission_id =?", id).Delete(&submission)
	if result.Error != nil {
		return "", result.Error
	}
	return "Submission deleted successfully", nil
}
