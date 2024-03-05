package main

import (
	"github.com/razorpay/GoLang-Learning/exercises"
	"github.com/razorpay/GoLang-Learning/exercises/salary"
)

func main() {

	// Basic exercises
	exercises.RunMatrix()
	exercises.MakeExpressionTree()
	salary.SalaryCalculation()

	//GoRoutine Related exercises
	exercises.CountLettersUsingGoRoutines()
	exercises.AverageRating()
	exercises.BankAccount()
}
