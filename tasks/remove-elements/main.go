package main

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/In-place_algorithm

func removeElement(nums []int, val int) int {
	count := 0
	for index := range nums {
		if nums[index] != val {
			nums[count] = nums[index]
			count++
		}
	}

	return count
}

func main() {
	// CASE1
	var v1 = []int{3, 2, 2, 3}
	var v2 = 3
	fmt.Print(removeElement(v1, v2)) // 2 [2,2]

	// CASE 2
	var v3 = []int{0, 1, 2, 2, 3, 0, 4, 2}
	var v4 = 2
	fmt.Print(removeElement(v3, v4)) // 5 [0,1,4,0,3]
}
