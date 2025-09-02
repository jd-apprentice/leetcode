package main

import (
	"fmt"
	"sort"
)

// https://www.geeksforgeeks.org/go-language/arrays-in-go/
// https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go
// https://www.geeksforgeeks.org/go-language/slices-in-golang/
// https://pkg.go.dev/golang.org/x/tools/internal/typesinternal#WrongResultCount
// https://pkg.go.dev/golang.org/x/tools/internal/typesinternal#TooManyValues

func merge(nums1 []int, m int, nums2 []int, n int) (response []int) {
	var slicedArr1 = nums1[0:m]
	var slicedArr2 = nums2[0:n]
	fullArr := append(slicedArr1, slicedArr2...)
	arrSlice := fullArr[:]
	sort.Ints(arrSlice)
	return arrSlice
}

func main() {
	// CASE 1
	v1 := []int{1, 2, 3, 0, 0, 0}
	v3 := []int{2, 5, 6}
	var v2 = 3
	var v4 = 3
	fmt.Print(merge(v1, v2, v3, v4))
}
