package models

import (
	"backend/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Suspended bool      `json:"suspended"`
}

func (student *Student) Create(email string, args ...string) (*Student, error) {
	student.Email = email
	if len(args) > 0 {
		student.Name = args[0]
	}

	if err := config.DB.Create(&student).Error; err != nil {
		return nil, err
	}

	return student, nil
}

func (student *Student) GetAll() ([]Student, error) {
	var students []Student
	if err := config.DB.Find(&student).Error; err != nil {
		return nil, err
	}

	return students, nil
}

func (student *Student) GetStudentByID() (*Student, error) {
	if err := config.DB.First(&student, student.ID).Error; err != nil {
		return nil, err
	}

	return student, nil
}

func (student *Student) UpdateStudentByEmail() error {
	// Update with conditions

	if err := config.DB.Model(&student).Where("email = ?", student.Email).Updates(map[string]interface{}{
		"name":      student.Name,
		"email":     student.Email,
		"suspended": student.Suspended,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (student *Student) DeleteStudentByID() error {
	if err := config.DB.Where("id = ?", student.ID).Delete(&student).Error; err != nil {
		return err
	}

	return nil
}
