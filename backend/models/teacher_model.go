package models

import (
	"backend/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func (teacher *Teacher) GetAll() ([]Teacher, error) {
	var teachers []Teacher
	if err := config.DB.Find(&teachers).Error; err != nil {
		return nil, err
	}

	return teachers, nil
}
