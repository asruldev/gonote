package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model           // Adds some metadata fields to the table
	ID         uuid.UUID `gorm:"type:char(36)"` // Explicitly specify the type to be uuid
	Title      string
	SubTitle   string
	Text       string
}

// This functions are called before creating any Post
// func (post *Note) BeforeCreate(scope *gorm.Scope) error {
// 	return scope.SetColumn("ID", uuid.NewV4())
//  }
