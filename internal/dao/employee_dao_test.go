package dao

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"test-assignment/internal/model"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	schema := `
	CREATE TABLE employees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		job_title TEXT NOT NULL,
		salary INTEGER NOT NULL
	);`

	if _, err := db.Exec(schema); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	return db
}

func TestEmployeeRepository_Create_Success(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewEmployeeRepository(db)

	emp := &model.Employee{
		FirstName: "Mahi",
		LastName:  "Kumar",
		JobTitle:  "Developer",
		Salary:    70000,
	}

	err := repo.Create(emp)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestEmployeeRepository_Create_DBError(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewEmployeeRepository(db)

	emp := &model.Employee{
		FirstName: "Mahi",
		LastName:  "Kumar",
		JobTitle:  "Developer",
		Salary:    70000,
	}

	err = repo.Create(emp)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
