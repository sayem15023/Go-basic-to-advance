package main

import "fmt"

func Average(numbers []int) float64 {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	avg := float64(sum) / float64(len(numbers))

	return avg

}

func main() {
	numbers := []int{10, 23, 34, 54}
	result := Average(numbers)
	fmt.Println("The avg is ", result)

}
