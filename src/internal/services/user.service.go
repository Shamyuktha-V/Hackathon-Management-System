package services

import (
	"Hackathon-Management-System/src/graph/model"
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/repositories"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	AppConfig      *appConfig.AppConfig
	UserRepository repositories.UserRepository
}

func NewUserService(appConfig *appConfig.AppConfig) *UserService {
	return &UserService{
		AppConfig:      appConfig,
		UserRepository: *repositories.NewUserRepository(appConfig),
	}
}

type IUserService interface {
	GetUser(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, input model.UpdateUserInput) (*model.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (string, error)
}

func (s UserService) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
	fmt.Print("User ID :: ", id)
	user, err := s.UserRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := s.UserRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s UserService) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	user, err := s.UserRepository.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s UserService) UpdateUser(ctx context.Context, id uuid.UUID, input model.UpdateUserInput) (*model.User, error) {
	user, err := s.UserRepository.UpdateUser(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s UserService) DeleteUser(ctx context.Context, id uuid.UUID) (string, error) {
	message, err := s.UserRepository.DeleteUser(ctx, id)
	if err != nil {
		return "", err
	}
	return message, nil
}
