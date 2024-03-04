package salary

import "fmt"

type IEmployeeTypeHandler interface {
	CalculateSalary() float64
}

type fullTimeEmployee struct {
	MonthlySalary float64
}
type contractor struct {
	DailyRate  float64
	DaysWorked int
}
type freelancer struct {
	HourlyRate  float64
	HoursWorked int
}

func (f fullTimeEmployee) CalculateSalary() float64 {
	return f.MonthlySalary
}

func (c contractor) CalculateSalary() float64 {
	return c.DailyRate * float64(c.DaysWorked)
}

func (fl freelancer) CalculateSalary() float64 {
	return fl.HourlyRate * float64(fl.HoursWorked)
}

func EmployeeTypeHandlerFactory(employeeType string) (IEmployeeTypeHandler, error) {
	switch employeeType {
	case FullTimeEmployee:
		return fullTimeEmployee{}, nil
	case Contractor:
		return contractor{}, nil
	case Freelancer:
		return freelancer{}, nil
	}

	return nil, fmt.Errorf("type of employee not exists")
}
