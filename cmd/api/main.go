package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mahigadamsetty/test-assignment/internal/dao"
	"github.com/mahigadamsetty/test-assignment/internal/handler"
	"github.com/mahigadamsetty/test-assignment/internal/service"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./employees.db")
	if err != nil {
		return nil, err
	}

	// Create employees table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS employees (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			full_name TEXT NOT NULL,
			job_title TEXT NOT NULL,
			country TEXT NOT NULL,
			salary REAL NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// Initialize database
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize layers
	employeeDAO := dao.NewEmployeeDAO(db)
	employeeService := service.NewEmployeeService(employeeDAO)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	// Setup routes
	router := mux.NewRouter()
	router.HandleFunc("/employees", employeeHandler.GetAllEmployees).Methods(http.MethodGet)
	router.HandleFunc("/employees", employeeHandler.CreateEmployee).Methods(http.MethodPost)
	router.HandleFunc("/employees/{id}", employeeHandler.GetEmployee).Methods(http.MethodGet)
	router.HandleFunc("/employees/{id}", employeeHandler.UpdateEmployee).Methods(http.MethodPut)
	router.HandleFunc("/employees/{id}", employeeHandler.DeleteEmployee).Methods(http.MethodDelete)

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
