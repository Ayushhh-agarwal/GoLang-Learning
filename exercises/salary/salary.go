package salary

import "fmt"

// 	Question :
//	There are 3 types of employees in an organization.,
//		FullTime
//		Contractor
//		Freelancer
//
//	Full-time employees are paid by the day
//	Contractors are paid by the day.
//	Freelancers are paid by the hour.
//
//	Example:
//		Full-time employee gets paid 15000 per month
//		The contractor gets paid 3000 per month
//		Freelancer gets paid 2000 (if freelancer works 20 hours)
//	Define an interface that calculates the salary of an employee.

func SalaryCalculation() {
	fullTimeEmpHandler, _ := EmployeeTypeHandlerFactory("fullTimeEmployee", 0, 0, 15000)
	fmt.Println("Full-time employee salary:", fullTimeEmpHandler.CalculateSalary())

	contractorHandler, _ := EmployeeTypeHandlerFactory("contractor", 20, 0, 3000)
	fmt.Println("Contractor salary:", contractorHandler.CalculateSalary())

	freelancerHandler, _ := EmployeeTypeHandlerFactory("freelancer", 0, 20, 100)
	fmt.Println("Freelancer salary:", freelancerHandler.CalculateSalary())
}
