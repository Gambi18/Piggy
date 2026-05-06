package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"piggy.com/internal/models"
	"piggy.com/internal/piggyservice"
)

type Handler struct {
	service *piggyservice.Service
}

func NewHandler(service *piggyservice.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateTransaction(c *gin.Context) {
	var payload models.CreateTransactionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.service.CreateTransaction(c, payload)
	if err != nil {
		errorMessage := fmt.Sprintf("Transaction failed for user %s (%s of %d): %v", payload.UserID, payload.Type, payload.Amount, err)
		fmt.Println(errorMessage)

		if strings.Contains(err.Error(), "insufficient balance") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Insufficient balance",
				"details": err.Error(),
				"context": errorMessage,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal server error",
			"details": err.Error(),
			"context": errorMessage,
		})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *Handler) GetTransactions(c *gin.Context) {
	type queryParams struct {
		UserID string `form:"userId"`
		Type   string `form:"type"`
	}
	var query queryParams

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var transactions *[]models.Transaction
	var err error

	if query.Type != "" {
		transactions, err = h.service.GetTransactionsByType(c, query.UserID, query.Type)
	} else {
		transactions, err = h.service.GetTransactions(c, query.UserID)
	}

	if err != nil {
		fmt.Printf("GetTransactions failed for user %s: %v\n", query.UserID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) GetUserBalance(c *gin.Context) {
	type queryParams struct {
		Username string `form:"username"`
	}
	var query queryParams

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.GetUserByUsername(c, query.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user.Balance)
}

func (h *Handler) SignUp(c *gin.Context) {
	var payload models.SignUpPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.SignUp(c, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) Login(c *gin.Context) {
	var payload models.SignInPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Login(c, payload)
	if err != nil {
		fmt.Printf("Login failed for user %s: %v\n", payload.Username, err)
		if err.Error() == "invalid password" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		// Assuming service returns "user not found" or similar for missing users
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetBalance(c *gin.Context) {
	type queryParams struct {
		UserID string `form:"userId"`
	}
	var query queryParams

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance, err := h.service.GetBalance(c, query.UserID)
	if err != nil {
		fmt.Printf("Error getting balance for user %s: %v\n", query.UserID, err)
		if err.Error() == "user not found" || err.Error() == "invalid user id format" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, balance)
}
