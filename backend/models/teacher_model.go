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

func (teacher *Teacher) Create() (*Teacher, error) {
	if err := config.DB.Create(&teacher).Error; err != nil {
		return nil, err
	}

	return teacher, nil
}

func (teacher *Teacher) GetAll() ([]Teacher, error) {
	var teachers []Teacher
	if err := config.DB.Find(&teachers).Error; err != nil {
		return nil, err
	}

	return teachers, nil
}

func (teacher *Teacher) GetTeacherByEmail() (*Teacher, error) {
	if err := config.DB.First(&teacher, teacher.Email).Error; err != nil {
		return nil, err
	}

	return teacher, nil
}

func (teacher *Teacher) UpdateTeacherByEmail() error {
	if err := config.DB.Model(&teacher).Where("email = ?", teacher.Email).Updates(map[string]interface{}{
		"name":  teacher.Name,
		"email": teacher.Email,
	}).Error; err != nil {
		return err
	}

	return nil
}

// Delete teacher by Email
func (teacher *Teacher) DeleteTeacherByEmail() error {
	if err := config.DB.Unscoped().Where("email = ?", teacher.Email).Delete(&teacher).Error; err != nil {
		return err
	}

	return nil
}
