package handlers

import (
	"net/http"

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}


func (h *Handler) GetTransactions(c *gin.Context) {
	transactions, err := h.service.GetTransactions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
