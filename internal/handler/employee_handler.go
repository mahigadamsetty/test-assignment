package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mahigadamsetty/test-assignment/internal/model"
)

// EmployeeService defines the interface for employee business logic operations
type EmployeeService interface {
	CreateEmployee(employee *model.Employee) error
	GetEmployeeByID(id int) (*model.Employee, error)
	GetAllEmployees() ([]*model.Employee, error)
	UpdateEmployee(employee *model.Employee) error
	DeleteEmployee(id int) error
}

// EmployeeHandler handles HTTP requests for employee endpoints
type EmployeeHandler struct {
	service EmployeeService
}

// NewEmployeeHandler creates a new EmployeeHandler instance
func NewEmployeeHandler(service EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

// CreateEmployee handles POST /employees
func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateEmployee(&employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
}

// GetEmployee handles GET /employees/{id}
func (h *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		// Fallback to manual parsing for tests
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(pathParts) >= 2 {
			idStr = pathParts[len(pathParts)-1]
		} else {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	employee, err := h.service.GetEmployeeByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

// GetAllEmployees handles GET /employees
func (h *EmployeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.service.GetAllEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// UpdateEmployee handles PUT /employees/{id}
func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		// Fallback to manual parsing for tests
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(pathParts) >= 2 {
			idStr = pathParts[len(pathParts)-1]
		} else {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var employee model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	employee.ID = id
	if err := h.service.UpdateEmployee(&employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

// DeleteEmployee handles DELETE /employees/{id}
func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		// Fallback to manual parsing for tests
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(pathParts) >= 2 {
			idStr = pathParts[len(pathParts)-1]
		} else {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteEmployee(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
