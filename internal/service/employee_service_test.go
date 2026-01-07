package service

import (
	"test-assignment/internal/model"
	"testing"
)

func Test_EmployeeService_CreateEmployee_Success(t *testing.T) {
	service := NewEmployeeService()

	employee := model.Employee{
		FirstName: "Mahi",
		LastName:  "Kumar",
		JobTitle:  "Developer",
		Salary:    70000,
	}

	err := service.CreateEmployee(employee)
	if err != nil {
		t.Errorf("CreateEmployee() error = %v, want nil", err)
	}
}
