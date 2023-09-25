package models

import (
	"backend/config"
)

type Registration struct {
	StudentEmail string `json:"student_email"`
	TeacherEmail string `json:"teacher_email"`
}

func (registration *Registration) Create(student_email string, teacher_email string) (*Registration, error) {
	registration.StudentEmail = student_email
	registration.TeacherEmail = teacher_email

	if err := config.DB.Create(&registration).Error; err != nil {
		return nil, err
	}

	return registration, nil
}

func (registration *Registration) GetAll() ([]Registration, error) {
	var registrations []Registration
	if err := config.DB.Find(&registrations).Error; err != nil {
		return nil, err
	}

	return registrations, nil
}

func (registration *Registration) GetRegistrationByTeacherEmail() ([]Registration, error) {
	var registrations []Registration

	// given a list of teachers, find students who are registered to ALL given teachers
	if err := config.DB.Where("teacher_email IN ?", registration.TeacherEmail).Find(&registrations).Error; err != nil {
		return nil, err
	}

	return registrations, nil
}
