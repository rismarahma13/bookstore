package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{"data": data})
}

func NotFoundResponse(c *gin.Context) {
	ErrorResponse(c, http.StatusNotFound, "Resource not found")
}

func InternalServerErrorResponse(c *gin.Context, err error) {
	ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Internal server error: %v", err))
}
