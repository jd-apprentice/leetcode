package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	trimedString := strings.Trim(s, " ")
	arrWords := strings.Split(trimedString, " ")
	return len(arrWords[len(arrWords)-1])
}

func main() {

	v1 := "Hello World"
	fmt.Print(lengthOfLastWord(v1))

	v2 := "   fly me   to   the moon  "
	fmt.Print(lengthOfLastWord(v2))
}
