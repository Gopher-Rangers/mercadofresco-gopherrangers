package handlers

import (
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/employee"
)

const (
	URL = "/api/v1/employees/"
)

type response struct {
	Code  int
	Data  []employee.Employee
	Error string
}

func createEmployeeArray() []employee.Employee {
	var emps []employee.Employee
	employee1 := employee.Employee{
		ID:          1,
		CardNumber:  117899,
		FirstName:   "Jose",
		LastName:    "Neves",
		WareHouseID: 456521,
	}
	employee2 := employee.Employee{
		ID:          2,
		CardNumber:  7878447,
		FirstName:   "Antonio",
		LastName:    "Moraes",
		WareHouseID: 11224411,
	}
	emps = append(emps, employee1, employee2)
	return emps
}
