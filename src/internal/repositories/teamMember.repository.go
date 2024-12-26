package repositories

import (
	"Hackathon-Management-System/src/graph/model"
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamMemberRepository struct {
	AppConfig *appConfig.AppConfig
	DB        *gorm.DB
}

func NewTeamMemberRepository(appConfig *appConfig.AppConfig) *TeamMemberRepository {
	return &TeamMemberRepository{
		AppConfig: appConfig,
		DB:        appConfig.DB,
	}
}

func InputToEntityTeamMember(input interface{}) models.TeamMember {
	switch input := input.(type) {
	case model.CreateTeamMemberInput:
		var entity models.TeamMember
		entity.ID = uuid.New()
		entity.TeamID, _ = uuid.Parse(input.TeamID)
		entity.UserID, _ = uuid.Parse(input.UserID)
		return entity

	case model.UpdateTeamMemberInput:
		var entity models.TeamMember
		if input.TeamID != nil {
			entity.TeamID, _ = uuid.Parse(*input.TeamID)
		}
		if input.UserID != nil {
			entity.UserID, _ = uuid.Parse(*input.UserID)
		}
		return entity
	}
	return models.TeamMember{}
}

func (repo *TeamMemberRepository) GetTeamMember(ctx context.Context, id uuid.UUID) (*model.TeamMember, error) {
	var teamMember *model.TeamMember

	result := repo.DB.Table(models.TeamMember{}.TableName()).Where("id =?", id).First(&teamMember)
	if result.Error != nil {
		return teamMember, result.Error
	}
	return teamMember, nil
}

func (repo *TeamMemberRepository) GetTeamMembersByTeamID(ctx context.Context, teamID uuid.UUID) ([]*model.TeamMember, error) {
	var teamMembers []*model.TeamMember

	result := repo.DB.Table(models.TeamMember{}.TableName()).Where("team_id =?", teamID).Find(&teamMembers)
	if result.Error != nil {
		return teamMembers, result.Error
	}
	return teamMembers, nil
}

func (repo *TeamMemberRepository) GetTeamMembersByUserID(ctx context.Context, userID uuid.UUID) ([]*model.TeamMember, error) {
	var teamMembers []*model.TeamMember

	result := repo.DB.Table(models.TeamMember{}.TableName()).Where("user_id =?", userID).Find(&teamMembers)
	if result.Error != nil {
		return teamMembers, result.Error
	}
	return teamMembers, nil
}

func (repo *TeamMemberRepository) CreateTeamMember(ctx context.Context, input model.CreateTeamMemberInput) (*model.TeamMember, error) {
	teamMember := InputToEntityTeamMember(input)
	fmt.Println("Team Member :: ", teamMember)
	result := repo.DB.Table(models.TeamMember{}.TableName()).Create(&teamMember)
	if result.Error != nil {
		return nil, result.Error
	}

	return repo.GetTeamMember(ctx, teamMember.ID)
}

func (repo *TeamMemberRepository) UpdateTeamMember(ctx context.Context, id uuid.UUID, input model.UpdateTeamMemberInput) (*model.TeamMember, error) {
	teamMember := InputToEntityTeamMember(input)

	result := repo.DB.Table(models.TeamMember{}.TableName()).Where("id =?", id).Updates(teamMember)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.GetTeamMember(ctx, id)
}

func (repo *TeamMemberRepository) DeleteTeamMember(ctx context.Context, id uuid.UUID) (string, error) {
	var teamMember *model.TeamMember

	result := repo.DB.Table(models.TeamMember{}.TableName()).Where("id =?", id).Delete(&teamMember)
	if result.Error != nil {
		return "", result.Error
	}
	return "Team Member deleted successfully", nil
}
