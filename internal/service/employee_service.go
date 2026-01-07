package service

import "test-assignment/internal/model"

type EmployeeServiceInterface interface {
	CreateEmployee(employee model.Employee) error
}

type employeeService struct{}

func NewEmployeeService() *employeeService {
	return &employeeService{}
}

func (s *employeeService) CreateEmployee(employee model.Employee) error {
	return nil
}
