package models

import (
	"time"

	"github.com/google/uuid"
)

// Hackathon struct represents the hackathon table
type Hackathon struct {
	ID               uuid.UUID `gorm:"primaryKey;column:id"`
	JudgeID          uuid.UUID `gorm:"not null;column:judge_id"`
	Name             string    `gorm:"size:255;not null;column:name"`
	ProblemStatement string    `gorm:"size:1000;column:problem_statement"`
	StartDate        time.Time `gorm:"not null;column:start_date"`
	EndDate          time.Time `gorm:"not null;column:end_date"`
	Duration         int       `gorm:"not null;column:duration"`
	FromDate         time.Time `gorm:"not null;column:from_date"`
	ToDate           time.Time `gorm:"not null;column:to_date"`
	CategoryID       uuid.UUID `gorm:"not null;column:category_id"`

	Judge    User     `gorm:"foreignKey:JudgeID;references:ID"`
	Category Category `gorm:"foreignKey:CategoryID;references:ID"`
}

// TableName overrides the default table name
func (h Hackathon) TableName() string {
	return "hackathons"
}
