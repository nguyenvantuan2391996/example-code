package handler

import (
	"face-be-service/internal/domains/employee"
)

type Handler struct {
	employeeService *employee.Employee
}

func NewHandler(employeeService *employee.Employee) *Handler {
	return &Handler{
		employeeService: employeeService,
	}
}
