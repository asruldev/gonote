package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36)"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
