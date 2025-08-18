package main

import "fmt"

func sumArray(arr []int) int {
	n := len(arr)
	total := 0

	for i := 0; i < n; i++ {
		total = total + arr[i]
	}
	return total
}
func main() {
	arr := []int{2, 3, 45, 6}
	total := sumArray(arr)
	fmt.Println("sum of array", total)
}
