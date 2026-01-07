package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"test-assignment/internal/model"
	"test-assignment/internal/service"
	"testing"
)

// Fake service just to satisfy handler dependency

type mockEmployeeService struct {
	createFn func(model.Employee) error
}

func (m *mockEmployeeService) CreateEmployee(emp model.Employee) error {
	return m.createFn(emp)
}

func TestEmployeeHandler_CreateEmployee_Success(t *testing.T) {
	mockSvc := &mockEmployeeService{
		createFn: func(emp model.Employee) error {
			return nil
		},
	}

	handler := NewEmployeeHandler(mockSvc)

	reqBody := `{"firstName":"Mahi","lastName":"Kumar","jobTitle":"Developer","country":"USA", "salary":70000}`
	req := httptest.NewRequest("POST", "/employees", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler.CreateEmployee(rec, req)
	if rec.Code != 201 {
		t.Errorf("expected status 201, got %d", rec.Code)
	}
}

// ---- tests ----

func TestCreateEmployee_InvalidEmployee_Returns400(t *testing.T) {
	mockSvc := &mockEmployeeService{
		createFn: func(emp model.Employee) error {
			return service.ErrInvalidEmployee
		},
	}

	handler := NewEmployeeHandler(mockSvc)

	reqBody := []byte(`{
		"first_name": "",
		"last_name": "Kumar",
		"job_title": "Developer",
		"salary": 70000
	}`)

	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.CreateEmployee(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}
