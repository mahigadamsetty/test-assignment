package handler

import (
	"net/http/httptest"
	"strings"
	"testing"
)

// Fake service just to satisfy handler dependency
type FakeEmployeeService struct{}

func TestEmployeeHandler_CreateEmployee_Success(t *testing.T) {
	fakeService := &FakeEmployeeService{}

	handler := NewEmployeeHandler(fakeService)

	reqBody := `{"firstName":"Mahi","lastName":"Kumar","jobTitle":"Developer","country":"USA", "salary":70000}`
	req := httptest.NewRequest("POST", "/employees", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler.CreateEmployee(rec, req)
	if rec.Code != 201 {
		t.Errorf("expected status 201, got %d", rec.Code)
	}
}
