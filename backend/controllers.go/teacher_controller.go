package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type TeacherController struct{}

func (tc *TeacherController) CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get the teacher name from the request body
	teacher.ID = uuid.New()
	name := c.PostForm("name")

	// create a new teacher record with the specified name and email
	if _, err := teacher.Create(teacher.Email, name); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Teacher created successfully",
			"teacher": teacher,
		})
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (tc *TeacherController) GetAllTeachers(c *gin.Context) {
	var teacher models.Teacher
	teachers, err := teacher.GetAll()
	if err == nil {
		// create a new slice to hold the results
		results := make([]gin.H, len(teachers))
		for i, teacher := range teachers {
			// add only the id and email fields to the results slice
			results[i] = gin.H{
				"id":    teacher.ID,
				"email": teacher.Email,
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "All Teachers",
			"teacher": results,
		})
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (tc *TeacherController) UpdateTeacherByID(c *gin.Context) {
	var teacher models.Teacher
	id := c.Param("id")
	teacher.ID = uuid.MustParse(id)
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := teacher.UpdateTeacherByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teacher updated successfully"})
}

func (tc *TeacherController) DeleteTeacherByID(c *gin.Context) {
	var teacher models.Teacher
	id := c.Param("id")
	teacher.ID = uuid.MustParse(id)
	err := teacher.DeleteTeacherByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}
