package errors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error types for better error handling
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

// Common error constructors
func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewInternalServerErrorWithDetails(message, details string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
		Details: details,
	}
}

// Error response handler
func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.Code, gin.H{
			"error": appErr.Message,
			"code":  appErr.Code,
		})
		if appErr.Details != "" {
			c.JSON(appErr.Code, gin.H{
				"details": appErr.Details,
			})
		}
		return
	}

	// Handle unexpected errors
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal server error",
		"code":  http.StatusInternalServerError,
	})
}

// Validation error helper
func HandleValidationError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   "Validation failed",
		"details": err.Error(),
		"code":    http.StatusBadRequest,
	})
}

// Success response helper
func HandleSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

// Created response helper
func HandleCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    data,
	})
}

// Log error with context
func LogError(operation string, userID interface{}, err error) {
	fmt.Printf("ERROR in %s for user %v: %v\n", operation, userID, err)
}
