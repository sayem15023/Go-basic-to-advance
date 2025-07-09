package main

import "fmt"

func isPalindrom(str string) bool {
	charCount := make(map[rune]int)
	for _, ch := range str {
		charCount[ch]++
	}

	oddCount := 0

	for _, count := range charCount {
		if count%2 != 0 {
			oddCount++
		}
	}
	return oddCount <= 1
}

func main() {
	fmt.Println(isPalindrom("hello"))
}
