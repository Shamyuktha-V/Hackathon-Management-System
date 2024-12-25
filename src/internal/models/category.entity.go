package models

import (
	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `gorm:"primaryKey;column:id"`
	Name string    `gorm:"size:255;not null;column:name"`
}

// TableName overrides the default table name
func (c Category) TableName() string {
	return "categories"
}
