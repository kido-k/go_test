package main

import (
	"fmt"
)

func copySliceOne() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{10, 11}

	n := copy(s1, s2)

	fmt.Println("n= ", n)
	fmt.Println("s1= ", s1)
	fmt.Println("s2= ", s2)
}

func copySliceTwo() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{10, 11, 12, 13, 14, 15}

	n := copy(s1, s2)

	fmt.Println("n= ", n)
	fmt.Println("s1= ", s1)
	fmt.Println("s2= ", s2)
}

func copySliceThree() {
	b := make([]byte, 9)

	n := copy(b, "あいうえお")

	fmt.Println("n= ", n)
	fmt.Println("b= ", b)
}

func main() {
	copySliceOne()
	copySliceTwo()
	copySliceThree()
}
