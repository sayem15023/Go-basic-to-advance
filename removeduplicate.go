package main

import "fmt"

func main() {
	numbers := []int{1, 2, 2, 5, 7}

	unique := []int{}

	seen := make(map[int]bool)

	for _, num := range numbers {
		if !seen[num] {
			unique = append(unique, num)
			seen[num] = true
		}
	}
	fmt.Println(unique)
}
