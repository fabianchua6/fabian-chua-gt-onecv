package main

import (
	"backend/config"
	"backend/controllers.go"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.DatabaseSetup()
	if err != nil {
		log.Println("Driver creation failed", err.Error())
	}

	//migrations.run()

	router := gin.Default()
	var teacherController controllers.TeacherController
	var registrationController controllers.RegistrationController
	// var studentController controllers.StudentController

	// teacher routes
	router.GET("/api/teachers", teacherController.GetAllTeachers)
	router.POST("/api/teachers", teacherController.CreateTeacher)
	router.PATCH("/api/teachers/:id", teacherController.UpdateTeacherByID)
	router.DELETE("/api/teachers/:id", teacherController.DeleteTeacherByID)

	// student routes
	// !TODO: implement student routes

	// registration routes
	router.GET("/api/register", registrationController.GetAll)
	router.POST("/api/register", registrationController.Create)
	// router.GET("/api/register/:teacher_email", registrationController.GetByTeacherEmail)

	router.Run("localhost:8080")
}
