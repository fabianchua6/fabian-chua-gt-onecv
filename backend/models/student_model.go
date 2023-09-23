package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
