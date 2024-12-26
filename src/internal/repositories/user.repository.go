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

type UserRepository struct {
	AppConfig *appConfig.AppConfig
	DB        *gorm.DB
}

func NewUserRepository(appConfig *appConfig.AppConfig) *UserRepository {
	return &UserRepository{
		AppConfig: appConfig,
		DB:        appConfig.DB,
	}
}

func InputToEntityUser(input interface{}) models.User {
	switch input := input.(type) {
	case model.CreateUserInput:
		var entity models.User
		entity.ID = uuid.New()
		entity.Name = input.Name
		entity.Email = input.Email
		if input.Role != nil {
			inputRole := models.RoleEnum(*input.Role)
			entity.Role = inputRole
		}
		return entity

	case model.UpdateUserInput:
		var entity models.User
		if input.Name != nil {
			entity.Name = *input.Name
		}
		if input.Email != nil {
			entity.Email = *input.Email
		}
		if input.Role != nil {
			inputRole := models.RoleEnum(*input.Role)
			entity.Role = inputRole
		}
		return entity
	}
	return models.User{}
}

func (repo *UserRepository) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var user *model.User
	fmt.Println("HELLo")
	result := repo.DB.Table(models.User{}.TableName()).Where("id =?", id).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	fmt.Println("User :", user)
	return user, nil
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user *model.User

	result := repo.DB.Table(models.User{}.TableName()).Where("email =?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (repo *UserRepository) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	user := InputToEntityUser(input)

	result := repo.DB.Table(models.User{}.TableName()).Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return repo.GetUser(ctx, user.ID)
}

func (repo *UserRepository) UpdateUser(ctx context.Context, id uuid.UUID, input model.UpdateUserInput) (*model.User, error) {
	user := InputToEntityUser(input)

	result := repo.DB.Table(models.User{}.TableName()).Where("id =?", id).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.GetUser(ctx, id)
}

func (repo *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) (string, error) {
	var user *model.User

	result := repo.DB.Table(models.User{}.TableName()).Where("id =?", id).Delete(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return "User deleted successfully", nil
}
