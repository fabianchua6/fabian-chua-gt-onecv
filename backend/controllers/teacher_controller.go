package controllers

import (
	"backend/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type TeacherController struct{}

type NotificationRequest struct {
	Teacher         string `json:"teacher" binding:"required"`
	NotificationMsg string `json:"notification" binding:"required"`
}

func (tc *TeacherController) CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		StandardErrorResponse((c), http.StatusBadRequest, err.Error())
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

func (tc *TeacherController) UpdateTeacherByEmail(c *gin.Context) {
	var teacher models.Teacher
	teacher.Email = c.Param("email")
	if err := c.ShouldBindJSON(&teacher); err != nil {
		StandardErrorResponse((c), http.StatusBadRequest, err.Error())
		return
	}
	err := teacher.UpdateTeacherByEmail()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teacher updated successfully"})
}

func (tc *TeacherController) DeleteTeacherByEmail(c *gin.Context) {
	var teacher models.Teacher
	teacher.Email = c.Param("email")

	err := teacher.DeleteTeacherByEmail()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}

// send notification
// criteria
// 1. provided they are not suspended
// 2. all students who are registered to the teacher
// 3. OR they are @ in the notification

func (tc *TeacherController) GetNotifiedStudents(c *gin.Context) {
	var notificationRequest NotificationRequest

	// take teacher email and notification from the following body
	if err := c.ShouldBindJSON(&notificationRequest); err != nil {

		StandardErrorResponse((c), http.StatusBadRequest, err.Error())

		return
	}

	// get the teacher email from the request body
	teacherEmail := notificationRequest.Teacher

	// get the students tagged from NotificationMsg from the request body
	notificationMsg := notificationRequest.NotificationMsg

	// Define the regex pattern for an email
	// from https://emailregex.com/
	regexPattern := `@[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`

	// Compile the regex
	re := regexp.MustCompile(regexPattern)

	// Find all instances of the pattern in the given text
	emailMatches := re.FindAllString(notificationMsg, -1)

	// Create or Initialize the notifiedStudents list
	notifiedStudents := make([]string, len(emailMatches))

	// Append each match to the notifiedStudents list without the "@" symbol
	for i, email := range emailMatches {
		notifiedStudents[i] = email[1:]
	}

	// add those who are registered to this teacher
	var registration models.Registration
	registrations, _ := registration.GetRegistrationByTeacherEmail([]string{teacherEmail})

	// from registration get the student emails add add to notifiedStudents
	for _, registration := range registrations {
		notifiedStudents = append(notifiedStudents, registration.StudentEmail)
	}

	// from the notifiedStudents check they are not suspended
	// remove suspended students from notifiedStudents

	var finalResultSet []string

	var student models.Student
	for i := 0; i < len(notifiedStudents); i++ {
		student.Email = notifiedStudents[i]
		suspended, _ := student.IsSuspended()

		if !suspended {
			// add to new resultSet
			finalResultSet = append(finalResultSet, notifiedStudents[i])
		}
	}

	// return list of notified students
	c.JSON(http.StatusOK, gin.H{
		"recipients": finalResultSet,
	})
}
