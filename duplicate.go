package main

import "fmt"

func duplicateRemove(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}

	}
	return i + 1

}

func main() {
	num1 := []int{1, 2, 3, 1, 2}
	k1 := duplicateRemove(num1)
	fmt.Println("K", k1, "nums", num1[:k1])
}
