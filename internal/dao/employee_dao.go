package dao

import (
	"database/sql"

	"github.com/mahigadamsetty/test-assignment/internal/model"
)

// EmployeeDAO handles database operations for employees
type EmployeeDAO struct {
	db *sql.DB
}

// NewEmployeeDAO creates a new EmployeeDAO instance
func NewEmployeeDAO(db *sql.DB) *EmployeeDAO {
	return &EmployeeDAO{db: db}
}

// Create inserts a new employee into the database
func (dao *EmployeeDAO) Create(employee *model.Employee) error {
	result, err := dao.db.Exec(`
		INSERT INTO employees (full_name, job_title, country, salary)
		VALUES (?, ?, ?, ?)
	`, employee.FullName, employee.JobTitle, employee.Country, employee.Salary)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	employee.ID = int(id)
	return nil
}

// GetByID retrieves an employee by their ID
func (dao *EmployeeDAO) GetByID(id int) (*model.Employee, error) {
	employee := &model.Employee{}
	err := dao.db.QueryRow(`
		SELECT id, full_name, job_title, country, salary
		FROM employees
		WHERE id = ?
	`, id).Scan(&employee.ID, &employee.FullName, &employee.JobTitle, &employee.Country, &employee.Salary)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

// GetAll retrieves all employees from the database
func (dao *EmployeeDAO) GetAll() ([]*model.Employee, error) {
	rows, err := dao.db.Query(`
		SELECT id, full_name, job_title, country, salary
		FROM employees
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*model.Employee
	for rows.Next() {
		employee := &model.Employee{}
		err := rows.Scan(&employee.ID, &employee.FullName, &employee.JobTitle, &employee.Country, &employee.Salary)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, rows.Err()
}

// Update updates an existing employee in the database
func (dao *EmployeeDAO) Update(employee *model.Employee) error {
	_, err := dao.db.Exec(`
		UPDATE employees
		SET full_name = ?, job_title = ?, country = ?, salary = ?
		WHERE id = ?
	`, employee.FullName, employee.JobTitle, employee.Country, employee.Salary, employee.ID)

	return err
}

// Delete removes an employee from the database
func (dao *EmployeeDAO) Delete(id int) error {
	_, err := dao.db.Exec(`
		DELETE FROM employees
		WHERE id = ?
	`, id)

	return err
}
