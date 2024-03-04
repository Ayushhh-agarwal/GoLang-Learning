package salary

import "fmt"

func SalaryCalculation() {
	fullTimeEmp := fullTimeEmployee{MonthlySalary: 15000}
	contractors := contractor{DailyRate: 3000, DaysWorked: 20}
	freelancers := freelancer{HourlyRate: 100, HoursWorked: 20}

	fmt.Println("Full-time employee salary:", fullTimeEmp.CalculateSalary())
	fmt.Println("Contractor salary:", contractors.CalculateSalary())
	fmt.Println("Freelancer salary:", freelancers.CalculateSalary())

	//TODO: Make this handlerFactory work as expected
	empHandler, _ := EmployeeTypeHandlerFactory("freelancer")
	fmt.Println("Freelancer salary:", empHandler.CalculateSalary())
}
