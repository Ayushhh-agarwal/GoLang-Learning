package salary

import "fmt"

func SalaryCalculation() {
	fullTimeEmpHandler, _ := EmployeeTypeHandlerFactory("fullTimeEmployee", 0, 0, 15000)
	fmt.Println("Full-time employee salary:", fullTimeEmpHandler.CalculateSalary())

	contractorHandler, _ := EmployeeTypeHandlerFactory("contractor", 20, 0, 3000)
	fmt.Println("Contractor salary:", contractorHandler.CalculateSalary())

	freelancerHandler, _ := EmployeeTypeHandlerFactory("freelancer", 0, 20, 100)
	fmt.Println("Freelancer salary:", freelancerHandler.CalculateSalary())
}
