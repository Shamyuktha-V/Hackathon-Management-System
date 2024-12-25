package models

import "github.com/google/uuid"

type TeamMember struct {
	ID     uuid.UUID gorm:"primaryKey;column:id"     // Primary key for the team member
	TeamID uuid.UUID gorm:"not null;column:team_id"  // Foreign key referencing Team table
	UserID uuid.UUID gorm:"not null;column:user_id"  // Foreign key referencing User table
}

func (tm TeamMember) TableName() string {
	return "team_members"
}
