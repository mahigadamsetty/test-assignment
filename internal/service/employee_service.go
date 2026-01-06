package service

import (
	"errors"

	"github.com/mahigadamsetty/test-assignment/internal/model"
)

// EmployeeDAO defines the interface for employee data access operations
type EmployeeDAO interface {
	Create(employee *model.Employee) error
	GetByID(id int) (*model.Employee, error)
	GetAll() ([]*model.Employee, error)
	Update(employee *model.Employee) error
	Delete(id int) error
}

// EmployeeService handles business logic for employees
type EmployeeService struct {
	dao EmployeeDAO
}

// NewEmployeeService creates a new EmployeeService instance
func NewEmployeeService(dao EmployeeDAO) *EmployeeService {
	return &EmployeeService{dao: dao}
}

// validateEmployee validates all required fields and business rules
func (s *EmployeeService) validateEmployee(employee *model.Employee) error {
	if employee.FullName == "" {
		return errors.New("full name is required")
	}
	if employee.JobTitle == "" {
		return errors.New("job title is required")
	}
	if employee.Country == "" {
		return errors.New("country is required")
	}
	if employee.Salary <= 0 {
		return errors.New("salary must be positive")
	}
	return nil
}

// CreateEmployee creates a new employee with validation
func (s *EmployeeService) CreateEmployee(employee *model.Employee) error {
	if err := s.validateEmployee(employee); err != nil {
		return err
	}
	return s.dao.Create(employee)
}

// GetEmployeeByID retrieves an employee by ID
func (s *EmployeeService) GetEmployeeByID(id int) (*model.Employee, error) {
	return s.dao.GetByID(id)
}

// GetAllEmployees retrieves all employees
func (s *EmployeeService) GetAllEmployees() ([]*model.Employee, error) {
	return s.dao.GetAll()
}

// UpdateEmployee updates an existing employee with validation
func (s *EmployeeService) UpdateEmployee(employee *model.Employee) error {
	if err := s.validateEmployee(employee); err != nil {
		return err
	}
	return s.dao.Update(employee)
}

// DeleteEmployee deletes an employee by ID
func (s *EmployeeService) DeleteEmployee(id int) error {
	return s.dao.Delete(id)
}
