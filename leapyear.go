package main

import "fmt"

func main() {
	var year int64
	fmt.Println("Input the year")
	fmt.Scanln(&year)
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println(year, "Year is leap year")
	} else {
		fmt.Println(year, "year is not a leap year ")
	}

}
