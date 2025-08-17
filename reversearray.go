package main

import "fmt"

// reverse function takes an integer slice and returns the reversed slice
func reverse(arr []int) []int {
	n := len(arr)
	newArr := make([]int, n)
	for i := 0; i < n; i++ {
		newArr[i] = arr[n-1-i]
		// swap elements at position i and n-1-i
		//arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return newArr
}

func main() {
	// Original array
	arr := []int{1, 2, 3, 4}

	// Call reverse function
	reversed := reverse(arr)

	// Print result
	fmt.Println("Original Array:", arr)
	fmt.Println("Reversed Array:", reversed)
}
