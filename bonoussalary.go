package main

import "fmt"

func BonusSalary(baseSalary float64, overTimeHours int) float64 {

	var bonous = 1.10

	total := baseSalary

	if overTimeHours > 10 {
		total *= bonous
	}
	return float64(total)

}

func main() {
	salary := BonusSalary(10000, 20)
	fmt.Println("Final salary of this employee is ", salary)
}
