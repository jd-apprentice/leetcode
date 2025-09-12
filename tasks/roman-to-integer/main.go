package main

import (
	"fmt"
)

func romanToInt(s string) int {
	romanNumbers := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prevVal := 0

	for index := len(s) - 1; index >= 0; index-- {
		currentVal := romanNumbers[s[index]]
		multiplier := 1
		if currentVal < prevVal {
			multiplier = -1
		}
		total += currentVal * multiplier
		prevVal = currentVal
	}

	return total
}

func main() {
	v1 := "III" // 3
	fmt.Printf("%s = %d\n", v1, romanToInt(v1))

	v2 := "LVIII" // 58
	fmt.Printf("%s = %d\n", v2, romanToInt(v2))

	v3 := "MCMXCIV" // 1994
	fmt.Printf("%s = %d\n", v3, romanToInt(v3))
}
