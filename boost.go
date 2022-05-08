package main

import (
	"go-microservices/Controllers/Employees"
)

func testAppEngine() {
	//TODO: here we have to ascertain about the object creational of the whole framework
	var engin = Employees.EmployeeList{}
	var params map[string]string = map[string]string{"context": "service"}
	//engin.Run(params)
	engin.RunMe(&engin, params)
}
