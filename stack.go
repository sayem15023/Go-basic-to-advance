package main

import "fmt"

func main() {
	stack := []string{}

	stack = append(stack, "Mango")
	stack = append(stack, "Banana")
	stack = append(stack, "Jackfruit")

	fmt.Println("Stack after pushing", stack)

	for len(stack) > 0 {
		n := len(stack) - 1
		fruit := stack[n]
		stack = stack[:n]
		fmt.Println("poped", fruit)
	}
}
