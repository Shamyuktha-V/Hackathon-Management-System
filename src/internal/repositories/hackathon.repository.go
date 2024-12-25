package repositories

import (
	"Hackathon-Management-System/src/graph/model"
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/models"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HackathonRepository struct {
	AppConfig *appConfig.AppConfig
	DB        *gorm.DB
}

func NewHackathonRepository(appConfig *appConfig.AppConfig) *HackathonRepository {
	return &HackathonRepository{
		AppConfig: appConfig,
		DB:        appConfig.DB,
	}
}

func InputToEntityHackathon(input interface{}) models.Hackathon {
	switch input := input.(type) {
	case model.CreateHackathonInput:
		var entity models.Hackathon
		entity.JudgeID, _ = uuid.Parse(input.JudgeID)
		entity.Name = input.Name
		entity.ProblemStatement = input.ProblemStatement
		entity.StartDate = input.StartDate
		entity.EndDate = input.EndDate
		entity.Duration = input.Duration
		entity.FromDate = input.FromDate
		entity.ToDate = input.ToDate
		entity.CategoryID, _ = uuid.Parse(input.CategoryID)
		return entity

	case model.UpdateHackathonInput:
		var entity models.Hackathon
		if input.JudgeID != nil {
			entity.JudgeID, _ = uuid.Parse(*input.JudgeID)
		}
		if input.Name != nil {
			entity.Name = *input.Name
		}
		if input.ProblemStatement != nil {
			entity.ProblemStatement = *input.ProblemStatement
		}
		if input.StartDate != nil {
			entity.StartDate = *input.StartDate
		}
		if input.EndDate != nil {
			entity.EndDate = *input.EndDate
		}
		if input.Duration != nil {
			entity.Duration = *input.Duration
		}
		if input.FromDate != nil {
			entity.FromDate = *input.FromDate
		}
		if input.ToDate != nil {
			entity.ToDate = *input.ToDate
		}
		if input.CategoryID != nil {
			entity.CategoryID, _ = uuid.Parse(*input.CategoryID)
		}
		return entity
	}
	return models.Hackathon{}
}

// GetHackathon fetches a hackathon by ID.
func (repo *HackathonRepository) GetHackathon(ctx context.Context, id uuid.UUID) (*model.Hackathon, error) {
	var hackathon *model.Hackathon
	result := repo.DB.Table(models.Hackathon{}.TableName()).Where("id =?", id).First(&hackathon)
	if result.Error != nil {
		return hackathon, result.Error
	}
	return hackathon, nil
}

// GetHackathonByCategory fetches a hackathon by category ID.
func (repo *HackathonRepository) GetHackathonByCategory(ctx context.Context, categoryID uuid.UUID) (*model.Hackathon, error) {
	var hackathon *model.Hackathon
	result := repo.DB.Table(models.Hackathon{}.TableName()).Where("category_id =?", categoryID).First(&hackathon)
	if result.Error != nil {
		return hackathon, result.Error
	}
	return hackathon, nil
}

// GetHackathonsByAttributes fetches hackathons by multiple attributes.
func (repo *HackathonRepository) GetHackathonsByAttributes(ctx context.Context, judgeID *uuid.UUID, name *string, problemStatement *string, startDate *string, endDate *string, categoryID *uuid.UUID, duration *int) ([]*model.Hackathon, error) {
	var hackathons []*model.Hackathon
	query := repo.DB.Table(models.Hackathon{}.TableName())

	// Dynamically applying filters based on non-nil parameters.
	if judgeID != nil {
		query = query.Where("judge_id = ?", *judgeID)
	}
	if name != nil {
		query = query.Where("name = ?", *name)
	}
	if problemStatement != nil {
		query = query.Where("problem_statement = ?", *problemStatement)
	}
	if startDate != nil {
		query = query.Where("start_date = ?", *startDate)
	}
	if endDate != nil {
		query = query.Where("end_date = ?", *endDate)
	}
	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}
	if duration != nil {
		query = query.Where("duration = ?", *duration)
	}

	// Execute the query
	result := query.Find(&hackathons)
	if result.Error != nil {
		return nil, result.Error
	}
	return hackathons, nil
}

// CreateHackathon creates a new hackathon.
func (repo *HackathonRepository) CreateHackathon(ctx context.Context, input model.CreateHackathonInput) (*model.Hackathon, error) {
	hackathon := InputToEntityHackathon(input)

	result := repo.DB.Table(models.Hackathon{}.TableName()).Create(&hackathon)
	if result.Error != nil {
		return nil, result.Error
	}

	return repo.GetHackathon(ctx, hackathon.ID)
}

// UpdateHackathon updates an existing hackathon.
func (repo *HackathonRepository) UpdateHackathon(ctx context.Context, id uuid.UUID, input model.UpdateHackathonInput) (*model.Hackathon, error) {
	hackathon := InputToEntityHackathon(input)
	query := fmt.Sprintf("UPDATE %s SET version = version + 1 WHERE id =?", models.Hackathon{}.TableName())

	if err := repo.DB.Exec(query, id).Error; err != nil {
		return nil, err
	}

	result := repo.DB.Table(models.Hackathon{}.TableName()).Where("id =?", id).Updates(hackathon)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.GetHackathon(ctx, id)
}

// DeleteHackathon deletes an existing hackathon.
func (repo *HackathonRepository) DeleteHackathon(ctx context.Context, id uuid.UUID) (string, error) {
	var hackathon *model.Hackathon

	result := repo.DB.Table(models.Hackathon{}.TableName()).Where("id =?", id).Delete(&hackathon)
	if result.Error != nil {
		return "", result.Error
	}
	return "Hackathon deleted successfully", nil
}

// GetHackathonsAvailableForRegistration fetches hackathons that are currently available for registration.
func (repo *HackathonRepository) GetHackathonsAvailableForRegistration(ctx context.Context) ([]*model.Hackathon, error) {
	var hackathons []*model.Hackathon
	currentDate := fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05")) // Format current date in the same format as the database

	// Query to fetch hackathons where the current date is between startDate and endDate
	result := repo.DB.Table(models.Hackathon{}.TableName()).
		Where("start_date <= ? AND end_date >= ?", currentDate, currentDate).
		Find(&hackathons)
	if result.Error != nil {
		return nil, result.Error
	}
	return hackathons, nil
}

// GetFutureHackathons fetches hackathons with a startDate greater than the current date.
func (repo *HackathonRepository) GetFutureHackathons(ctx context.Context) ([]*model.Hackathon, error) {
	var hackathons []*model.Hackathon
	currentDate := fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05")) // Format current date in the same format as the database

	// Query to fetch hackathons with a startDate greater than the current date
	result := repo.DB.Table(models.Hackathon{}.TableName()).
		Where("start_date > ?", currentDate).
		Find(&hackathons)
	if result.Error != nil {
		return nil, result.Error
	}
	return hackathons, nil
}
