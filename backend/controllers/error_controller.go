package controllers

import "github.com/gin-gonic/gin"

func StandardErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"message": message,
	})
}
