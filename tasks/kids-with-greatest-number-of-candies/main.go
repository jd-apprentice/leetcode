package main

import (
	"fmt"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {

	result := []bool{}

	maxValue := findLargestNumber(candies)
	for index := range len(candies) {
		if candies[index]+extraCandies >= maxValue {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}

func findLargestNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	largest := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}
	}
	return largest
}

func main() {

	v1 := []int{2, 3, 5, 1, 3} // [true,true,true,false,true]
	v2 := 3

	fmt.Print(kidsWithCandies(v1, v2))
}
