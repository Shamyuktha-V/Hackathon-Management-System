package services

import (
	"Hackathon-Management-System/src/graph/model"
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/repositories"
	"context"

	"github.com/google/uuid"
)

type CategoryService struct {
	AppConfig          *appConfig.AppConfig
	CategoryRepository repositories.CategoryRepository
}

func NewCategoryService(appConfig *appConfig.AppConfig) *CategoryService {
	return &CategoryService{
		AppConfig:          appConfig,
		CategoryRepository: *repositories.NewCategoryRepository(appConfig),
	}
}

type ICategoryService interface {
	GetCategory(ctx context.Context, id uuid.UUID) (*model.Category, error)
	GetAllCategories(ctx context.Context) ([]*model.Category, error)
	GetCategoryByName(ctx context.Context, name string) (*model.Category, error)
	CreateCategory(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, input model.UpdateCategoryInput) (*model.Category, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) (string, error)
}

func (s CategoryService) GetCategory(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	category, err := s.CategoryRepository.GetCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s CategoryService) GetCategoryByName(ctx context.Context, name string) (*model.Category, error) {
	category, err := s.CategoryRepository.GetCategoryByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s CategoryService) GetAllCategories(ctx context.Context) ([]*model.Category, error) {
	categories, err := s.CategoryRepository.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s CategoryService) CreateCategory(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error) {
	category, err := s.CategoryRepository.CreateCategory(ctx, input)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s CategoryService) UpdateCategory(ctx context.Context, id uuid.UUID, input model.UpdateCategoryInput) (*model.Category, error) {
	category, err := s.CategoryRepository.UpdateCategory(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s CategoryService) DeleteCategory(ctx context.Context, id uuid.UUID) (string, error) {
	message, err := s.CategoryRepository.DeleteCategory(ctx, id)
	if err != nil {
		return "", err
	}
	return message, nil
}
