package dao

import (
	"database/sql"

	"test-assignment/internal/model"
)

type EmployeeRepositoryInterface interface {
	Create(emp *model.Employee) error
}

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepositoryInterface {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(emp *model.Employee) error {
	query := `
	INSERT INTO employees (first_name, last_name, job_title, salary)
	VALUES (?, ?, ?, ?)
	`

	result, err := r.db.Exec(
		query,
		emp.FirstName,
		emp.LastName,
		emp.JobTitle,
		emp.Salary,
	)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
