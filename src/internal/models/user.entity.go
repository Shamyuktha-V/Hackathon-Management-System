package models

import "github.com/google/uuid"

type User struct {
	ID      uuid.UUID `gorm:"size:36;primaryKey;column:id"`
	Name    string    `gorm:"size:255;not null;column:name"`
	Email   string    `gorm:"size:255;unique;not null;column:email"`
	Role    RoleEnum  `gorm:"type:enum('admin','judge','participant');not null;default:'participant';column:role"`
	Version int       `gorm:"default:1;column:version"`
}
type RoleEnum string

const (
	RoleAdmin       RoleEnum = "admin"
	RoleParticipant RoleEnum = "participant"
	RoleJudge       RoleEnum = "judge"
)

func (u User) TableName() string {
	return "users"
}
