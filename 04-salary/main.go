package main

import "fmt"

func main() {
	var salary float64
	var minSalary float64 = 8_000_000
	var taxPercent float64 = 0

	fmt.Println("enter your salary: ")
	fmt.Scan(&salary)

	if salary <= minSalary {
		taxPercent = 0
	} else if salary > minSalary && salary <= minSalary*2 {
		taxPercent = 0.05
	} else if salary > minSalary*2 && salary <= minSalary*3 {
		taxPercent = 0.07
	} else if salary > minSalary*3 && salary <= minSalary*4 {
		taxPercent = 0.10
	} else {
		taxPercent = 0.15
	}

	fmt.Printf("your tax is: %.2f\n", taxPercent*salary)
	fmt.Printf("your salary is: %.2f\n", salary-taxPercent*salary)
}
