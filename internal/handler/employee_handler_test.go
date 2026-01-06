package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mahigadamsetty/test-assignment/internal/model"
)

// MockService is a mock implementation of EmployeeService for testing
type MockService struct {
	CreateEmployeeFunc      func(*model.Employee) error
	GetEmployeeByIDFunc     func(int) (*model.Employee, error)
	GetAllEmployeesFunc     func() ([]*model.Employee, error)
	UpdateEmployeeFunc      func(*model.Employee) error
	DeleteEmployeeFunc      func(int) error
}

func (m *MockService) CreateEmployee(e *model.Employee) error {
	if m.CreateEmployeeFunc != nil {
		return m.CreateEmployeeFunc(e)
	}
	e.ID = 1
	return nil
}

func (m *MockService) GetEmployeeByID(id int) (*model.Employee, error) {
	if m.GetEmployeeByIDFunc != nil {
		return m.GetEmployeeByIDFunc(id)
	}
	return &model.Employee{ID: id, FullName: "Test", JobTitle: "Test", Country: "Test", Salary: 50000}, nil
}

func (m *MockService) GetAllEmployees() ([]*model.Employee, error) {
	if m.GetAllEmployeesFunc != nil {
		return m.GetAllEmployeesFunc()
	}
	return []*model.Employee{}, nil
}

func (m *MockService) UpdateEmployee(e *model.Employee) error {
	if m.UpdateEmployeeFunc != nil {
		return m.UpdateEmployeeFunc(e)
	}
	return nil
}

func (m *MockService) DeleteEmployee(id int) error {
	if m.DeleteEmployeeFunc != nil {
		return m.DeleteEmployeeFunc(id)
	}
	return nil
}

func TestEmployeeHandler_CreateEmployee_Success(t *testing.T) {
	mockService := &MockService{}
	handler := NewEmployeeHandler(mockService)

	employee := map[string]interface{}{
		"full_name": "John Doe",
		"job_title": "Engineer",
		"country":   "USA",
		"salary":    75000.00,
	}

	body, _ := json.Marshal(employee)
	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.CreateEmployee(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("CreateEmployee() status = %v, want %v", rec.Code, http.StatusCreated)
	}

	var result model.Employee
	json.NewDecoder(rec.Body).Decode(&result)

	if result.ID == 0 {
		t.Error("CreateEmployee() did not return employee with ID")
	}
}

func TestEmployeeHandler_CreateEmployee_InvalidJSON(t *testing.T) {
	mockService := &MockService{}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.CreateEmployee(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("CreateEmployee() status = %v, want %v", rec.Code, http.StatusBadRequest)
	}
}

func TestEmployeeHandler_CreateEmployee_ValidationError(t *testing.T) {
	mockService := &MockService{
		CreateEmployeeFunc: func(e *model.Employee) error {
			return errors.New("full name is required")
		},
	}
	handler := NewEmployeeHandler(mockService)

	employee := map[string]interface{}{
		"full_name": "",
		"job_title": "Engineer",
		"country":   "USA",
		"salary":    75000.00,
	}

	body, _ := json.Marshal(employee)
	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.CreateEmployee(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("CreateEmployee() status = %v, want %v", rec.Code, http.StatusBadRequest)
	}
}

func TestEmployeeHandler_GetEmployee_Success(t *testing.T) {
	mockService := &MockService{
		GetEmployeeByIDFunc: func(id int) (*model.Employee, error) {
			return &model.Employee{
				ID:       id,
				FullName: "Jane Smith",
				JobTitle: "Manager",
				Country:  "UK",
				Salary:   85000.00,
			}, nil
		},
	}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/employees/1", nil)
	rec := httptest.NewRecorder()

	handler.GetEmployee(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("GetEmployee() status = %v, want %v", rec.Code, http.StatusOK)
	}

	var result model.Employee
	json.NewDecoder(rec.Body).Decode(&result)

	if result.FullName != "Jane Smith" {
		t.Errorf("GetEmployee() FullName = %v, want Jane Smith", result.FullName)
	}
}

func TestEmployeeHandler_GetEmployee_NotFound(t *testing.T) {
	mockService := &MockService{
		GetEmployeeByIDFunc: func(id int) (*model.Employee, error) {
			return nil, sql.ErrNoRows
		},
	}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/employees/999", nil)
	rec := httptest.NewRecorder()

	handler.GetEmployee(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("GetEmployee() status = %v, want %v", rec.Code, http.StatusNotFound)
	}
}

func TestEmployeeHandler_GetEmployee_InvalidID(t *testing.T) {
	mockService := &MockService{}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/employees/abc", nil)
	rec := httptest.NewRecorder()

	handler.GetEmployee(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("GetEmployee() status = %v, want %v", rec.Code, http.StatusBadRequest)
	}
}

func TestEmployeeHandler_GetAllEmployees(t *testing.T) {
	mockService := &MockService{
		GetAllEmployeesFunc: func() ([]*model.Employee, error) {
			return []*model.Employee{
				{ID: 1, FullName: "Alice", JobTitle: "Designer", Country: "Canada", Salary: 65000},
				{ID: 2, FullName: "Bob", JobTitle: "Developer", Country: "Australia", Salary: 70000},
			}, nil
		},
	}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/employees", nil)
	rec := httptest.NewRecorder()

	handler.GetAllEmployees(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("GetAllEmployees() status = %v, want %v", rec.Code, http.StatusOK)
	}

	var result []*model.Employee
	json.NewDecoder(rec.Body).Decode(&result)

	if len(result) != 2 {
		t.Errorf("GetAllEmployees() returned %d employees, want 2", len(result))
	}
}

func TestEmployeeHandler_UpdateEmployee_Success(t *testing.T) {
	mockService := &MockService{}
	handler := NewEmployeeHandler(mockService)

	employee := map[string]interface{}{
		"full_name": "Tom Brown",
		"job_title": "Senior Analyst",
		"country":   "Germany",
		"salary":    70000.00,
	}

	body, _ := json.Marshal(employee)
	req := httptest.NewRequest(http.MethodPut, "/employees/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.UpdateEmployee(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("UpdateEmployee() status = %v, want %v", rec.Code, http.StatusOK)
	}
}

func TestEmployeeHandler_UpdateEmployee_InvalidID(t *testing.T) {
	mockService := &MockService{}
	handler := NewEmployeeHandler(mockService)

	employee := map[string]interface{}{
		"full_name": "Tom Brown",
		"job_title": "Senior Analyst",
		"country":   "Germany",
		"salary":    70000.00,
	}

	body, _ := json.Marshal(employee)
	req := httptest.NewRequest(http.MethodPut, "/employees/abc", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.UpdateEmployee(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("UpdateEmployee() status = %v, want %v", rec.Code, http.StatusBadRequest)
	}
}

func TestEmployeeHandler_UpdateEmployee_ValidationError(t *testing.T) {
	mockService := &MockService{
		UpdateEmployeeFunc: func(e *model.Employee) error {
			return errors.New("full name is required")
		},
	}
	handler := NewEmployeeHandler(mockService)

	employee := map[string]interface{}{
		"full_name": "",
		"job_title": "Senior Analyst",
		"country":   "Germany",
		"salary":    70000.00,
	}

	body, _ := json.Marshal(employee)
	req := httptest.NewRequest(http.MethodPut, "/employees/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.UpdateEmployee(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("UpdateEmployee() status = %v, want %v", rec.Code, http.StatusBadRequest)
	}
}

func TestEmployeeHandler_DeleteEmployee_Success(t *testing.T) {
	mockService := &MockService{}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/employees/1", nil)
	rec := httptest.NewRecorder()

	handler.DeleteEmployee(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("DeleteEmployee() status = %v, want %v", rec.Code, http.StatusNoContent)
	}
}

func TestEmployeeHandler_DeleteEmployee_InvalidID(t *testing.T) {
	mockService := &MockService{}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/employees/abc", nil)
	rec := httptest.NewRecorder()

	handler.DeleteEmployee(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("DeleteEmployee() status = %v, want %v", rec.Code, http.StatusBadRequest)
	}
}

func TestEmployeeHandler_DeleteEmployee_ServiceError(t *testing.T) {
	mockService := &MockService{
		DeleteEmployeeFunc: func(id int) error {
			return errors.New("database error")
		},
	}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/employees/1", nil)
	rec := httptest.NewRecorder()

	handler.DeleteEmployee(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("DeleteEmployee() status = %v, want %v", rec.Code, http.StatusInternalServerError)
	}
}
