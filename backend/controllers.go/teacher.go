package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type TeacherController struct{}

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
