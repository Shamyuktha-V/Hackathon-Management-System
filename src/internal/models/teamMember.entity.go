package models

import "github.com/google/uuid"

type TeamMember struct {
	ID     uuid.UUID `gorm:"primaryKey;column:id"`            // Primary key for the team member
	TeamID uuid.UUID `gorm:"size:36;not null;column:team_id"` // Foreign key referencing Team table
	Team   Team      `gorm:"foreignKey:TeamID;references:ID"`
	UserID uuid.UUID `gorm:"size:36;not null;column:user_id"` // Foreign key referencing User table
}

func (tm TeamMember) TableName() string {
	return "team_members"
}
