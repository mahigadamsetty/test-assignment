package handler

import (
	"net/http"
)

type EmployeeService interface{}

type EmployeeHandler struct {
	service EmployeeService // Dependency injection of service
}

func NewEmployeeHandler(s EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: s}
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
