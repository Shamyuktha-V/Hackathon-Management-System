package repositories

import (
	"Hackathon-Management-System/src/graph/model"
	appConfig "Hackathon-Management-System/src/internal/config"
	"Hackathon-Management-System/src/internal/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	AppConfig *appConfig.AppConfig
	DB        *gorm.DB
}

func NewCategoryRepository(appConfig *appConfig.AppConfig) *CategoryRepository {
	return &CategoryRepository{
		AppConfig: appConfig,
		DB:        appConfig.DB,
	}
}

// InputToEntityCategory converts Create/Update Category input into the models.Category entity
func InputToEntityCategory(input interface{}) models.Category {
	switch input := input.(type) {
	case model.CreateCategoryInput:
		var entity models.Category
		entity.ID = uuid.New()
		entity.Name = input.Name
		return entity

	case model.UpdateCategoryInput:
		var entity models.Category
		if input.Name != nil {
			entity.Name = *input.Name
		}
		return entity
	}
	return models.Category{}
}

// GetCategory fetches a category by ID.
func (repo *CategoryRepository) GetCategory(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	var category *model.Category
	result := repo.DB.Table(models.Category{}.TableName()).Where("id =?", id).First(&category)
	if result.Error != nil {
		return category, result.Error
	}
	return category, nil
}

func (repo *CategoryRepository) GetCategoryByName(ctx context.Context, name string) (*model.Category, error) {
	var category *model.Category
	result := repo.DB.Table(models.Category{}.TableName()).Where("name =?", name).First(&category)
	if result.Error != nil {
		return category, result.Error
	}
	return category, nil
}

// GetCategories fetches all categories.
func (repo *CategoryRepository) GetAllCategories(ctx context.Context) ([]*model.Category, error) {
	var categories []*model.Category
	result := repo.DB.Table(models.Category{}.TableName()).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

// CreateCategory creates a new category.
func (repo *CategoryRepository) CreateCategory(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error) {
	category := InputToEntityCategory(input)

	result := repo.DB.Table(models.Category{}.TableName()).Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return repo.GetCategory(ctx, category.ID)
}

// UpdateCategory updates an existing category.
func (repo *CategoryRepository) UpdateCategory(ctx context.Context, id uuid.UUID, input model.UpdateCategoryInput) (*model.Category, error) {
	category := InputToEntityCategory(input)

	// Update category in the database
	result := repo.DB.Table(models.Category{}.TableName()).Where("id =?", id).Updates(category)
	if result.Error != nil {
		return nil, result.Error
	}

	return repo.GetCategory(ctx, id)
}

// DeleteCategory deletes an existing category.
func (repo *CategoryRepository) DeleteCategory(ctx context.Context, id uuid.UUID) (string, error) {
	var category *model.Category
	result := repo.DB.Table(models.Category{}.TableName()).Where("id =?", id).Delete(&category)
	if result.Error != nil {
		return "", result.Error
	}
	return "Category deleted successfully", nil
}
