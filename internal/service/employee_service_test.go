package service

import (
	"testing"

	"test-assignment/internal/model"
)

type mockEmployeeRepository struct{}

func (m *mockEmployeeRepository) Create(emp *model.Employee) error {
	return nil
}

func TestEmployeeService_CreateEmployee_Validation(t *testing.T) {
	mockRepo := &mockEmployeeRepository{}
	service := NewEmployeeService(mockRepo)

	tests := []struct {
		name     string
		employee model.Employee
		wantErr  bool
	}{
		{
			name: "empty first name",
			employee: model.Employee{
				FirstName: "",
				LastName:  "Kumar",
				JobTitle:  "Developer",
				Salary:    70000,
			},
			wantErr: true,
		},
		{
			name: "empty last name",
			employee: model.Employee{
				FirstName: "Mahi",
				LastName:  "",
				JobTitle:  "Developer",
				Salary:    70000,
			},
			wantErr: true,
		},
		{
			name: "empty job title",
			employee: model.Employee{
				FirstName: "Mahi",
				LastName:  "Kumar",
				JobTitle:  "",
				Salary:    70000,
			},
			wantErr: true,
		},
		{
			name: "salary is zero",
			employee: model.Employee{
				FirstName: "Mahi",
				LastName:  "Kumar",
				JobTitle:  "Developer",
				Salary:    0,
			},
			wantErr: true,
		},
		{
			name: "salary is negative",
			employee: model.Employee{
				FirstName: "Mahi",
				LastName:  "Kumar",
				JobTitle:  "Developer",
				Salary:    -5000,
			},
			wantErr: true,
		},
		{
			name: "valid employee",
			employee: model.Employee{
				FirstName: "Mahi",
				LastName:  "Kumar",
				JobTitle:  "Developer",
				Salary:    70000,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.CreateEmployee(tt.employee)
			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}
