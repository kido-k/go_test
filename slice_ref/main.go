package main

import (
	"fmt"
)

func pow(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func pow2(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func main() {
	a := [3]int{1, 2, 3}
	pow(a)
	fmt.Println(a)

	a2 := []int{1, 2, 3}
	pow2(a2)
	fmt.Println(a2)
}
