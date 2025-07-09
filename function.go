package main

import (
	"fmt"
	"math"
)

func CheckAndSquareRoot(n float64) (float64, string) {
	if n > 0 {
		return math.Sqrt(n), "Positive"

	} else {
		return -1, "Negative"
	}
}

func main() {
	result, message := CheckAndSquareRoot(25)
	fmt.Println("Result", result, "|Message", message)
}
