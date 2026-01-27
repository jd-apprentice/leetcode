package main

import (
	"fmt"
)

// Example 1:

// Input: digits = [1,2,3]
// Output: [1,2,4]
// Explanation: The array represents the integer 123.
// Incrementing by one gives 123 + 1 = 124.
// Thus, the result should be [1,2,4].

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

func main() {

	v1 := []int{1, 2, 3}
	v2 := []int{4, 3, 2, 1}
	v3 := []int{9}

	fmt.Print(plusOne(v1)) // 1,2,4
	fmt.Print(plusOne(v2)) // 4,3,2,2
	fmt.Print(plusOne(v3)) // 1,0
}
