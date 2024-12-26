package repositories

import (
	"Hackathon-Management-System/src/graph/model"
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamRepository struct {
	AppConfig *appConfig.AppConfig
	DB        *gorm.DB
}

func NewTeamRepository(appConfig *appConfig.AppConfig) *TeamRepository {
	return &TeamRepository{
		AppConfig: appConfig,
		DB:        appConfig.DB,
	}
}

// InputToEntityTeam converts GraphQL input types to entity types
func InputToEntityTeam(input interface{}) models.Team {
	switch input := input.(type) {
	case model.CreateTeamInput:
		var entity models.Team
		entity.ID = uuid.New()
		entity.TeamName = input.TeamName
		entity.LeaderID, _ = uuid.Parse(input.LeaderID)
		entity.TeamSize = input.TeamSize
		return entity

	case model.UpdateTeamInput:
		var entity models.Team
		if input.TeamName != nil {
			entity.TeamName = *input.TeamName
		}
		if input.LeaderID != nil {
			entity.LeaderID, _ = uuid.Parse(*input.LeaderID)
		}
		if input.TeamSize != nil {
			entity.TeamSize = *input.TeamSize
		}
		return entity
	}
	return models.Team{}
}

// GetTeam retrieves a team by its ID
func (repo *TeamRepository) GetTeam(ctx context.Context, id uuid.UUID) (*model.Team, error) {
	var team *model.Team

	result := repo.DB.Table(models.Team{}.TableName()).Where("id =?", id).First(&team)
	if result.Error != nil {
		return team, result.Error
	}
	return team, nil
}

// GetTeamByLeader retrieves a team by its leader's ID
func (repo *TeamRepository) GetTeamByLeader(ctx context.Context, leaderID uuid.UUID) (*model.Team, error) {
	var team *model.Team

	result := repo.DB.Table(models.Team{}.TableName()).Where("leader_id =?", leaderID).First(&team)
	if result.Error != nil {
		return team, result.Error
	}
	return team, nil
}

// CreateTeam creates a new team
func (repo *TeamRepository) CreateTeam(ctx context.Context, input model.CreateTeamInput) (*model.Team, error) {
	team := InputToEntityTeam(input)

	result := repo.DB.Table(models.Team{}.TableName()).Create(&team)
	if result.Error != nil {
		return nil, result.Error
	}

	return repo.GetTeam(ctx, team.ID)
}

// UpdateTeam updates an existing team
func (repo *TeamRepository) UpdateTeam(ctx context.Context, id uuid.UUID, input model.UpdateTeamInput) (*model.Team, error) {
	team := InputToEntityTeam(input)

	result := repo.DB.Table(models.Team{}.TableName()).Where("id =?", id).Updates(team)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.GetTeam(ctx, id)
}

// DeleteTeam deletes a team by its ID
func (repo *TeamRepository) DeleteTeam(ctx context.Context, id uuid.UUID) (string, error) {
	var team *model.Team

	result := repo.DB.Table(models.Team{}.TableName()).Where("id =?", id).Delete(&team)
	if result.Error != nil {
		return "", result.Error
	}
	return "Team deleted successfully", nil
}
