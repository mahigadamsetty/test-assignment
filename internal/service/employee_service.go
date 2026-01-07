package service

import (
	"errors"
	"strings"

	"test-assignment/internal/dao"
	"test-assignment/internal/model"
)

var ErrInvalidEmployee = errors.New("invalid employee data")
var ErrEmployeeNotFound = errors.New("employee not found")

type EmployeeServiceInterface interface {
	CreateEmployee(employee model.Employee) error
}

type employeeService struct {
	repo dao.EmployeeRepositoryInterface
}

func NewEmployeeService(repo dao.EmployeeRepositoryInterface) *employeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) CreateEmployee(employee model.Employee) error {
	if strings.TrimSpace(employee.FirstName) == "" {
		return ErrInvalidEmployee
	}

	if strings.TrimSpace(employee.LastName) == "" {
		return ErrInvalidEmployee
	}

	if strings.TrimSpace(employee.JobTitle) == "" {
		return ErrInvalidEmployee
	}

	if employee.Salary <= 0 {
		return ErrInvalidEmployee
	}

	// DAO call will come later
	return nil
}
