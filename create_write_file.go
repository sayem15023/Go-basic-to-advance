package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("error", err)
	}
	defer file.Close()
	_, err = file.WriteString("Hello ,Go file")
	if err != nil {
		fmt.Println("error writing to file ", err)
		return
	}
	fmt.Println("file written properly ")

}
