package main

import "fmt"

func countdown(n int) {
	if n == 0 {
		fmt.Println("Done")
		return
	}
	fmt.Println(n)
	countdown(n - 1)
}

func main() {
	countdown(5)
}
