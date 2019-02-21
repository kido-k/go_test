package main

import (
	"fmt"
)

func plus(a, b int) []int {
	plus := a + b
	minus := a - b
	multi := a * b
	div := a / b
	rem := a % b
	array := []int{plus, minus, multi, div, rem}
	return array
}

func main() {
	var x, y int
	var result []int
	x = 2
	y = 3

	result = plus(x, y)
	fmt.Println(result)
}
