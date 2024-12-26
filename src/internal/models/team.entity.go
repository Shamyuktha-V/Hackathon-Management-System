package models

import "github.com/google/uuid"

type Team struct {
	ID       uuid.UUID `gorm:"primaryKey;column:id"`
	TeamName string    `gorm:"size:255;not null;column:team_name"`
	LeaderID uuid.UUID `gorm:"size:36;not null;column:leader_id"`
	Leader   User      `gorm:"foreignKey:LeaderID;references:ID"`
	TeamSize int       `gorm:"not null;column:team_size"`
}

func (t Team) TableName() string {
	return "teams"
}
