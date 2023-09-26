package main

import (
	"backend/config"
	"backend/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.DatabaseSetup()
	if err != nil {
		log.Println("Driver creation failed", err.Error())
	}

	router := gin.Default()
	var teacherController controllers.TeacherController
	var registrationController controllers.RegistrationController
	var studentController controllers.StudentController

	// teacher routes
	router.POST("/api/teachers", teacherController.CreateTeacher)
	router.GET("/api/teachers", teacherController.GetAllTeachers)
	router.DELETE("/api/teachers/:email", teacherController.DeleteTeacherByEmail)

	router.POST("/api/retrievefornotifications", teacherController.GetNotifiedStudents)

	// student routes
	router.POST("/api/students", studentController.CreateStudent)
	router.GET("/api/students", studentController.GetAllStudents)
	router.DELETE("/api/students/:email", studentController.DeleteStudentByEmail)

	router.POST("/api/suspend", studentController.SuspendStudentByEmail)

	// registration routes
	router.GET("/api/register", registrationController.GetAll)
	router.POST("/api/register", registrationController.Create)

	router.GET("/api/commonstudents", registrationController.GetByTeacherEmails)

	router.Run("localhost:8080")
}
