package controller

import "github.com/gin-gonic/gin"

func JsonMessageWithStatus(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"message": message,
	})
}

func JsonMessageWithStatusAndHeaders(c *gin.Context, statusCode int, message string, headers map[string]string) {
	for key, value := range headers {
		c.Header(key, value)
	}

	JsonMessageWithStatus(c, statusCode, message)
}

func AddHeaders(c *gin.Context, headers map[string]string) {
	for key, value := range headers {
		c.Header(key, value)
	}
}
