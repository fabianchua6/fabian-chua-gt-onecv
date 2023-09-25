package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StudentController struct{}

type SuspendRequest struct {
	Student string `json:"student" binding:"required"`
}

func (sc *StudentController) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get the student name from the request body
	student.ID = uuid.New()
	name := c.PostForm("name")

	// create a new student record with the specified name and email
	if _, err := student.Create(student.Email, name); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Student created successfully",
			"student": student,
		})
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (sc *StudentController) GetAllStudents(c *gin.Context) {
	var student models.Student
	students, err := student.GetAll()
	if err == nil {
		// create a new slice to hold the results
		results := make([]gin.H, len(students))
		for i, student := range students {
			// add only the id and email fields to the results slice
			results[i] = gin.H{
				"id":    student.ID,
				"email": student.Email,
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "All Students",
			"students": results,
		})
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (sc *StudentController) UpdateStudentByEmail(c *gin.Context) {
	var student models.Student
	student.Email = c.Param("email")
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := student.UpdateStudentByEmail()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}

func (sc *StudentController) SuspendStudentByEmail(c *gin.Context) {
	var student models.Student
	var suspendRequest SuspendRequest

	// take student email from the following body
	if err := c.ShouldBindJSON(&suspendRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student.Email = suspendRequest.Student
	student.Suspended = true

	if err := student.UpdateStudentByEmail(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (sc *StudentController) DeleteStudentByID(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	student.ID = uuid.MustParse(id)
	err := student.DeleteStudentByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
