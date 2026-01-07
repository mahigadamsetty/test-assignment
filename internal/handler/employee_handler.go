package handler

import (
	"errors"
	"net/http"
	"test-assignment/internal/model"
	"test-assignment/internal/service"

	"github.com/gin-gonic/gin"
)

type EmployeeService interface {
	CreateEmployee(employee model.Employee) error
}

type EmployeeHandler struct {
	service EmployeeService // Dependency injection of service
}

func NewEmployeeHandler(s EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: s}
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var emp model.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	err := h.service.CreateEmployee(emp)
	if err != nil {
		if errors.Is(err, service.ErrInvalidEmployee) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Employee created successfully"})
}
