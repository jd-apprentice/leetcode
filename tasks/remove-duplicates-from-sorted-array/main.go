package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	count := 1
	for index := range len(nums) {
		if index < 1 {
			continue
		}

		if nums[index-1] != nums[index] {
			nums[count] = nums[index]
			count++
		}
	}

	return count
}

func main() {
	// CASE 1
	var v1 = []int{1, 1, 2}
	fmt.Print(removeDuplicates(v1)) // 2 [1,2,_]

	fmt.Print("\n")

	// CASE 2
	var v2 = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Print(removeDuplicates(v2)) // 5 [0,1,2,3,4]
}
