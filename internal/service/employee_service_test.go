package service

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/mahigadamsetty/test-assignment/internal/model"
)

// MockDAO is a mock implementation of DAO for testing
type MockDAO struct {
	CreateFunc   func(*model.Employee) error
	GetByIDFunc  func(int) (*model.Employee, error)
	GetAllFunc   func() ([]*model.Employee, error)
	UpdateFunc   func(*model.Employee) error
	DeleteFunc   func(int) error
}

func (m *MockDAO) Create(e *model.Employee) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(e)
	}
	e.ID = 1
	return nil
}

func (m *MockDAO) GetByID(id int) (*model.Employee, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return &model.Employee{ID: id, FullName: "Test", JobTitle: "Test", Country: "Test", Salary: 50000}, nil
}

func (m *MockDAO) GetAll() ([]*model.Employee, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return []*model.Employee{}, nil
}

func (m *MockDAO) Update(e *model.Employee) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(e)
	}
	return nil
}

func (m *MockDAO) Delete(id int) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}

func TestEmployeeService_CreateEmployee_Success(t *testing.T) {
	mockDAO := &MockDAO{}
	service := NewEmployeeService(mockDAO)

	employee := &model.Employee{
		FullName: "John Doe",
		JobTitle: "Engineer",
		Country:  "USA",
		Salary:   75000.00,
	}

	err := service.CreateEmployee(employee)
	if err != nil {
		t.Errorf("CreateEmployee() error = %v, want nil", err)
	}

	if employee.ID == 0 {
		t.Error("CreateEmployee() did not set employee ID")
	}
}

func TestEmployeeService_CreateEmployee_MissingFullName(t *testing.T) {
	mockDAO := &MockDAO{}
	service := NewEmployeeService(mockDAO)

	employee := &model.Employee{
		FullName: "",
		JobTitle: "Engineer",
		Country:  "USA",
		Salary:   75000.00,
	}

	err := service.CreateEmployee(employee)
	if err == nil {
		t.Error("CreateEmployee() expected error for missing full name, got nil")
	}

	expectedError := "full name is required"
	if err.Error() != expectedError {
		t.Errorf("CreateEmployee() error = %v, want %v", err, expectedError)
	}
}

func TestEmployeeService_CreateEmployee_MissingJobTitle(t *testing.T) {
	mockDAO := &MockDAO{}
	service := NewEmployeeService(mockDAO)

	employee := &model.Employee{
		FullName: "John Doe",
		JobTitle: "",
		Country:  "USA",
		Salary:   75000.00,
	}

	err := service.CreateEmployee(employee)
	if err == nil {
		t.Error("CreateEmployee() expected error for missing job title, got nil")
	}

	expectedError := "job title is required"
	if err.Error() != expectedError {
		t.Errorf("CreateEmployee() error = %v, want %v", err, expectedError)
	}
}

func TestEmployeeService_CreateEmployee_MissingCountry(t *testing.T) {
	mockDAO := &MockDAO{}
	service := NewEmployeeService(mockDAO)

	employee := &model.Employee{
		FullName: "John Doe",
		JobTitle: "Engineer",
		Country:  "",
		Salary:   75000.00,
	}

	err := service.CreateEmployee(employee)
	if err == nil {
		t.Error("CreateEmployee() expected error for missing country, got nil")
	}

	expectedError := "country is required"
	if err.Error() != expectedError {
		t.Errorf("CreateEmployee() error = %v, want %v", err, expectedError)
	}
}

func TestEmployeeService_CreateEmployee_NegativeSalary(t *testing.T) {
	mockDAO := &MockDAO{}
	service := NewEmployeeService(mockDAO)

	employee := &model.Employee{
		FullName: "John Doe",
		JobTitle: "Engineer",
		Country:  "USA",
		Salary:   -1000.00,
	}

	err := service.CreateEmployee(employee)
	if err == nil {
		t.Error("CreateEmployee() expected error for negative salary, got nil")
	}

	expectedError := "salary must be positive"
	if err.Error() != expectedError {
		t.Errorf("CreateEmployee() error = %v, want %v", err, expectedError)
	}
}

func TestEmployeeService_CreateEmployee_ZeroSalary(t *testing.T) {
	mockDAO := &MockDAO{}
	service := NewEmployeeService(mockDAO)

	employee := &model.Employee{
		FullName: "John Doe",
		JobTitle: "Engineer",
		Country:  "USA",
		Salary:   0,
	}

	err := service.CreateEmployee(employee)
	if err == nil {
		t.Error("CreateEmployee() expected error for zero salary, got nil")
	}

	expectedError := "salary must be positive"
	if err.Error() != expectedError {
		t.Errorf("CreateEmployee() error = %v, want %v", err, expectedError)
	}
}

func TestEmployeeService_GetEmployeeByID_Success(t *testing.T) {
	mockDAO := &MockDAO{
		GetByIDFunc: func(id int) (*model.Employee, error) {
			return &model.Employee{
				ID:       id,
				FullName: "Jane Smith",
				JobTitle: "Manager",
				Country:  "UK",
				Salary:   85000.00,
			}, nil
		},
	}
	service := NewEmployeeService(mockDAO)

	employee, err := service.GetEmployeeByID(1)
	if err != nil {
		t.Errorf("GetEmployeeByID() error = %v, want nil", err)
	}

	if employee == nil {
		t.Fatal("GetEmployeeByID() returned nil employee")
	}

	if employee.FullName != "Jane Smith" {
		t.Errorf("GetEmployeeByID() FullName = %v, want Jane Smith", employee.FullName)
	}
}

func TestEmployeeService_GetEmployeeByID_NotFound(t *testing.T) {
	mockDAO := &MockDAO{
		GetByIDFunc: func(id int) (*model.Employee, error) {
			return nil, sql.ErrNoRows
		},
	}
	service := NewEmployeeService(mockDAO)

	employee, err := service.GetEmployeeByID(999)
	if err == nil {
		t.Error("GetEmployeeByID() expected error for not found, got nil")
	}

	if employee != nil {
		t.Error("GetEmployeeByID() should return nil employee when not found")
	}
}

func TestEmployeeService_GetAllEmployees(t *testing.T) {
	mockDAO := &MockDAO{
		GetAllFunc: func() ([]*model.Employee, error) {
			return []*model.Employee{
				{ID: 1, FullName: "Alice", JobTitle: "Designer", Country: "Canada", Salary: 65000},
				{ID: 2, FullName: "Bob", JobTitle: "Developer", Country: "Australia", Salary: 70000},
			}, nil
		},
	}
	service := NewEmployeeService(mockDAO)

	employees, err := service.GetAllEmployees()
	if err != nil {
		t.Errorf("GetAllEmployees() error = %v, want nil", err)
	}

	if len(employees) != 2 {
		t.Errorf("GetAllEmployees() returned %d employees, want 2", len(employees))
	}
}

func TestEmployeeService_UpdateEmployee_Success(t *testing.T) {
	mockDAO := &MockDAO{}
	service := NewEmployeeService(mockDAO)

	employee := &model.Employee{
		ID:       1,
		FullName: "Tom Brown",
		JobTitle: "Senior Analyst",
		Country:  "Germany",
		Salary:   70000.00,
	}

	err := service.UpdateEmployee(employee)
	if err != nil {
		t.Errorf("UpdateEmployee() error = %v, want nil", err)
	}
}

func TestEmployeeService_UpdateEmployee_ValidationFailure(t *testing.T) {
	mockDAO := &MockDAO{}
	service := NewEmployeeService(mockDAO)

	employee := &model.Employee{
		ID:       1,
		FullName: "",
		JobTitle: "Senior Analyst",
		Country:  "Germany",
		Salary:   70000.00,
	}

	err := service.UpdateEmployee(employee)
	if err == nil {
		t.Error("UpdateEmployee() expected validation error, got nil")
	}
}

func TestEmployeeService_DeleteEmployee_Success(t *testing.T) {
	mockDAO := &MockDAO{}
	service := NewEmployeeService(mockDAO)

	err := service.DeleteEmployee(1)
	if err != nil {
		t.Errorf("DeleteEmployee() error = %v, want nil", err)
	}
}

func TestEmployeeService_DeleteEmployee_DAOError(t *testing.T) {
	mockDAO := &MockDAO{
		DeleteFunc: func(id int) error {
			return errors.New("database error")
		},
	}
	service := NewEmployeeService(mockDAO)

	err := service.DeleteEmployee(1)
	if err == nil {
		t.Error("DeleteEmployee() expected error, got nil")
	}
}
