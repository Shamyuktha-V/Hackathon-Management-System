package models

import "github.com/google/uuid"

type Team struct {
	TeamID   uuid.UUID gorm:"primaryKey;column:team_id"
	TeamName string    gorm:"size:255;not null;column:team_name"
	LeaderID uuid.UUID gorm:"not null;column:leader_id" // Foreign key referencing User ID
	TeamSize int       gorm:"not null;column:team_size"
}

// TableName specifies the table name for the Team model
func (t Team) TableName() string {
	return "teams"
}
