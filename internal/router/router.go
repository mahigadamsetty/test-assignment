package router

import (
	"github.com/gin-gonic/gin"

	"test-assignment/internal/handler"
)

func NewRouter(employeeHandler *handler.EmployeeHandler) *gin.Engine {
	r := gin.New()

	// middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/employees", employeeHandler.CreateEmployee)
	}

	return r
}
