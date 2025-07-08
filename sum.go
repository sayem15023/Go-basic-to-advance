package main

import "fmt"

func Sum(x int64, y int64) int64 {
	result := x + y
	return result
}

func main() {
	result := Sum(10, 20)

	fmt.Println(result)
}
