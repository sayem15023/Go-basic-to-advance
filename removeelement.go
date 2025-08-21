package main

import "fmt"

func removeElements(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}

	}
	return k
}

func main() {
	nums := []int{1, 2, 3, 4, 6, 3}
	val := 2
	k := removeElements(nums, val)

	fmt.Println("the value of this code ", k, "value", nums[:k])

}
