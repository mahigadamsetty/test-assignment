package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"test-assignment/internal/model"
	"test-assignment/internal/service"
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

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp model.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.service.CreateEmployee(emp)
	if err != nil {
		if errors.Is(err, service.ErrInvalidEmployee) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
