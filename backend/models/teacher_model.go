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

func (teacher *Teacher) Create(email string, args ...string) (*Teacher, error) {
	teacher.Email = email
	if len(args) > 0 {
		teacher.Name = args[0]
	}

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

func (teacher *Teacher) GetTeacherByID() (*Teacher, error) {
	if err := config.DB.First(&teacher, teacher.ID).Error; err != nil {
		return nil, err
	}

	return teacher, nil
}

func (teacher *Teacher) UpdateTeacherByID() error {
	if err := config.DB.Model(&teacher).Where("id = ?", teacher.ID).Updates(map[string]interface{}{
		"name":  teacher.Name,
		"email": teacher.Email,
	}).Error; err != nil {
		return err
	}

	return nil
}

// Delete teacher by id
func (teacher *Teacher) DeleteTeacherByID() error {
	if err := config.DB.Where("id = ?", teacher.ID).Delete(&teacher).Error; err != nil {
		return err
	}

	return nil
}
