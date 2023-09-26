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

func (student *Student) Create() (*Student, error) {
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

// check if student is suspended by email
func (student *Student) IsSuspended() (bool, error) {
	var suspended bool
	if err := config.DB.Table("students").Select("suspended").Where("email = ?", student.Email).Scan(&suspended).Error; err != nil {
		return false, err
	}

	return suspended, nil
}

func (student *Student) GetStudentByEmail() (*Student, error) {
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

func (student *Student) DeleteStudentByEmail() error {
	if err := config.DB.Where("email = ?", student.Email).Delete(&student).Error; err != nil {
		return err
	}

	return nil
}
