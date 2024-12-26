package repositories

import (
	"Hackathon-Management-System/src/graph/model"
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubmissionRepository struct {
	AppConfig *appConfig.AppConfig
	DB        *gorm.DB
}

func NewSubmissionRepository(appConfig *appConfig.AppConfig) *SubmissionRepository {
	return &SubmissionRepository{
		AppConfig: appConfig,
		DB:        appConfig.DB,
	}
}

func InputToEntitySubmission(input interface{}) models.Submission {
	switch input := input.(type) {
	case model.CreateSubmissionInput:
		var entity models.Submission
		entity.ID = uuid.New()
		entity.TeamID, _ = uuid.Parse(input.TeamID)
		entity.HackathonID, _ = uuid.Parse(input.HackathonID)
		return entity

	case model.UpdateSubmissionInput:
		var entity models.Submission
		if input.TeamID != nil {
			entity.TeamID, _ = uuid.Parse(*input.TeamID)
		}
		if input.HackathonID != nil {
			entity.HackathonID, _ = uuid.Parse(*input.HackathonID)
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

func (repo *SubmissionRepository) GetSubmission(ctx context.Context, id uuid.UUID) (*model.Submission, error) {
	var submission *model.Submission

	result := repo.DB.Table(models.Submission{}.TableName()).Where("id =?", id).First(&submission)
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

func (repo *SubmissionRepository) GetHackathonsForAJudge(ctx context.Context, judgeID uuid.UUID) ([]*model.Submission, error) {
	var submission []*model.Submission

	result := repo.DB.Table(models.Submission{}.TableName()).Where("judge_id =?", judgeID).Find(&submission)
	if result.Error != nil {
		return submission, result.Error
	}
	return submission, nil
}

func (repo *SubmissionRepository) GetHackathonsForATeam(ctx context.Context, teamID uuid.UUID) ([]*model.Submission, error) {
	var submission []*model.Submission

	result := repo.DB.Table(models.Submission{}.TableName()).Preload("Team").Preload("Hackathon").Where("team_id =?", teamID).Find(&submission)
	if result.Error != nil {
		return submission, result.Error
	}
	return submission, nil
}

func (repo *SubmissionRepository) GetTeamsForAHackathons(ctx context.Context, hackathonID uuid.UUID) ([]*model.Submission, error) {
	var submission []*model.Submission

	result := repo.DB.Table(models.Submission{}.TableName()).Preload("Team").Preload("Hackathon").Where("hackathon_id =?", hackathonID).Find(&submission)
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

	return repo.GetSubmission(ctx, submission.ID)
}

func (repo *SubmissionRepository) UpdateSubmission(ctx context.Context, id uuid.UUID, input model.UpdateSubmissionInput) (*model.Submission, error) {
	submission := InputToEntitySubmission(input)

	result := repo.DB.Table(models.Submission{}.TableName()).Where("id =?", id).Updates(submission)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.GetSubmission(ctx, id)
}

func (repo *SubmissionRepository) DeleteSubmission(ctx context.Context, id uuid.UUID) (string, error) {
	var submission *model.Submission

	result := repo.DB.Table(models.Submission{}.TableName()).Where("id =?", id).Delete(&submission)
	if result.Error != nil {
		return "", result.Error
	}
	return "Submission deleted successfully", nil
}
