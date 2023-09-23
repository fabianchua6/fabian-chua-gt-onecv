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
	router.GET("/api/teachers", teacherController.GetAllTeachers)
	// router.POST("/api/teachers", handlers.CreateTeacher)
	router.Run("localhost:8080")
}
