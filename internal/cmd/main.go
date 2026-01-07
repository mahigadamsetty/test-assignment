package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"test-assignment/internal/dao"
	"test-assignment/internal/db"
	"test-assignment/internal/handler"
	"test-assignment/internal/router"
	"test-assignment/internal/service"
)

func main() {
	// ---- DB ----
	dbConn, err := sql.Open("sqlite3", "employees.db")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// ---- migrations ----
	if err := db.RunMigrations(dbConn, "internal/db/migrations"); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	// ---- DAO ----
	employeeRepo := dao.NewEmployeeRepository(dbConn)

	// ---- Service ----
	employeeService := service.NewEmployeeService(employeeRepo)

	// ---- Handler ----
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	// ---- Router ----
	r := router.NewRouter(employeeHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
