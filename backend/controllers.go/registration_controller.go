package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type RegistrationController struct{}

type RegistrationRequest struct {
	Teacher  string   `json:"teacher" binding:"required"`
	Students []string `json:"students" binding:"required"`
}

func (rc *RegistrationController) Create(c *gin.Context) {
	var registration models.Registration
	var registrationReq RegistrationRequest
	if err := c.ShouldBindJSON(&registrationReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create a new registration record with the specified teacher_email and loop through student_email
	for _, studentEmail := range registrationReq.Students {
		if _, err := registration.Create(studentEmail, registrationReq.Teacher); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	// return response http 204
	c.JSON(http.StatusNoContent, gin.H{})
}

func (rc *RegistrationController) GetAll(c *gin.Context) {
	var registration models.Registration
	registrations, err := registration.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message":       "All Registrations",
			"registrations": registrations,
		})
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

// func (rc *RegistrationController) GetByTeacherEmail(c *gin.Context) {
// 	var registration models.Registration

// 	// take in list of teacher emails separated by & in url parameters
// 	teacherEmails := strings.Split(c.Query("teacher_emails"), "&")

// 	teacherEmail := c.Param("teacher_email")
// 	registration.TeacherEmail = teacherEmail

// 	registrations, err := registration.GetRegistrationByTeacherEmail()

// 	// transform registration to slice of student emails
// 	var studentEmails []string
// 	for _, registration := range registrations {
// 		studentEmails = append(studentEmails, registration.StudentEmail)
// 	}

// 	if err == nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			// "message":  "Students registered to " + registration.TeacherEmail,
// 			"students": studentEmails,
// 		})
// 	} else {
// 		c.String(http.StatusInternalServerError, err.Error())
// 	}
// }

// func (tc *TeacherController) UpdateTeacherByID(c *gin.Context) {
// 	var teacher models.Teacher
// 	id := c.Param("id")
// 	teacher.ID = uuid.MustParse(id)
// 	if err := c.ShouldBindJSON(&teacher); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	err := teacher.UpdateTeacherByID()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Teacher updated successfully"})
// }

// func (tc *TeacherController) DeleteTeacherByID(c *gin.Context) {
// 	var teacher models.Teacher
// 	id := c.Param("id")
// 	teacher.ID = uuid.MustParse(id)
// 	err := teacher.DeleteTeacherByID()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
// }

// RegisterStudents to Teacher
