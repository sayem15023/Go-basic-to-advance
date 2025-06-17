package main

import "fmt"

func reversedString(input string) string {

	runes := []rune(input)

	i, j := 0, len(runes)-1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)

}

func main() {
	original := "mum"
	reversed := reversedString(original)
	if original == reversed {
		println("it is palindrome ")
	}
	fmt.Println("Original:", original)
	fmt.Println("Reversed:", reversed)
}
