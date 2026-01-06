package dao

import (
	"database/sql"
	"testing"

	"github.com/mahigadamsetty/test-assignment/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	// Create employees table
	_, err = db.Exec(`
		CREATE TABLE employees (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			full_name TEXT NOT NULL,
			job_title TEXT NOT NULL,
			country TEXT NOT NULL,
			salary REAL NOT NULL
		)
	`)
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	cleanup := func() {
		db.Close()
	}

	return db, cleanup
}

func TestEmployeeDAO_Create(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	dao := NewEmployeeDAO(db)

	employee := &model.Employee{
		FullName: "John Doe",
		JobTitle: "Software Engineer",
		Country:  "USA",
		Salary:   75000.00,
	}

	err := dao.Create(employee)
	if err != nil {
		t.Errorf("Create() error = %v, want nil", err)
	}

	if employee.ID == 0 {
		t.Error("Create() did not set employee ID")
	}
}

func TestEmployeeDAO_GetByID(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	dao := NewEmployeeDAO(db)

	// Insert test data
	result, err := db.Exec(`
		INSERT INTO employees (full_name, job_title, country, salary)
		VALUES (?, ?, ?, ?)
	`, "Jane Smith", "Product Manager", "UK", 85000.00)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	id, _ := result.LastInsertId()

	employee, err := dao.GetByID(int(id))
	if err != nil {
		t.Errorf("GetByID() error = %v, want nil", err)
	}

	if employee == nil {
		t.Fatal("GetByID() returned nil employee")
	}

	if employee.FullName != "Jane Smith" {
		t.Errorf("GetByID() FullName = %v, want Jane Smith", employee.FullName)
	}
}

func TestEmployeeDAO_GetByID_NotFound(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	dao := NewEmployeeDAO(db)

	employee, err := dao.GetByID(999)
	if err != sql.ErrNoRows {
		t.Errorf("GetByID() error = %v, want sql.ErrNoRows", err)
	}

	if employee != nil {
		t.Error("GetByID() should return nil employee when not found")
	}
}

func TestEmployeeDAO_GetAll(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	dao := NewEmployeeDAO(db)

	// Insert test data
	db.Exec(`INSERT INTO employees (full_name, job_title, country, salary) VALUES (?, ?, ?, ?)`,
		"Alice Johnson", "Designer", "Canada", 65000.00)
	db.Exec(`INSERT INTO employees (full_name, job_title, country, salary) VALUES (?, ?, ?, ?)`,
		"Bob Williams", "Developer", "Australia", 70000.00)

	employees, err := dao.GetAll()
	if err != nil {
		t.Errorf("GetAll() error = %v, want nil", err)
	}

	if len(employees) != 2 {
		t.Errorf("GetAll() returned %d employees, want 2", len(employees))
	}
}

func TestEmployeeDAO_Update(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	dao := NewEmployeeDAO(db)

	// Insert test data
	result, _ := db.Exec(`
		INSERT INTO employees (full_name, job_title, country, salary)
		VALUES (?, ?, ?, ?)
	`, "Tom Brown", "Analyst", "Germany", 60000.00)

	id, _ := result.LastInsertId()

	updatedEmployee := &model.Employee{
		ID:       int(id),
		FullName: "Tom Brown",
		JobTitle: "Senior Analyst",
		Country:  "Germany",
		Salary:   70000.00,
	}

	err := dao.Update(updatedEmployee)
	if err != nil {
		t.Errorf("Update() error = %v, want nil", err)
	}

	// Verify update
	employee, _ := dao.GetByID(int(id))
	if employee.JobTitle != "Senior Analyst" {
		t.Errorf("Update() JobTitle = %v, want Senior Analyst", employee.JobTitle)
	}
	if employee.Salary != 70000.00 {
		t.Errorf("Update() Salary = %v, want 70000.00", employee.Salary)
	}
}

func TestEmployeeDAO_Delete(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	dao := NewEmployeeDAO(db)

	// Insert test data
	result, _ := db.Exec(`
		INSERT INTO employees (full_name, job_title, country, salary)
		VALUES (?, ?, ?, ?)
	`, "Sam Davis", "Engineer", "France", 68000.00)

	id, _ := result.LastInsertId()

	err := dao.Delete(int(id))
	if err != nil {
		t.Errorf("Delete() error = %v, want nil", err)
	}

	// Verify deletion
	employee, err := dao.GetByID(int(id))
	if err != sql.ErrNoRows {
		t.Errorf("Delete() employee still exists, error = %v", err)
	}
	if employee != nil {
		t.Error("Delete() employee should be nil after deletion")
	}
}
