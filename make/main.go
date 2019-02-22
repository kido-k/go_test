package main

import (
	"fmt"
)

func createSlice() {
	// int型のスライス
	a := make([]int, 10)
	var b [10]int

	fmt.Println("a= ", a)
	fmt.Println("b= ", b)

	c := make([]float64, 3)
	fmt.Println("c= ", c)
	c[0] = 3.14
	c[1] = 6.28
	c[2] = 9.9999
	// c[3] = 9.9999 //エラー
	fmt.Println("c= ", c)
	fmt.Println("")
}

func checkSliceCapa() {
	s := make([]int, 5, 10)
	fmt.Println("s= ", s)
	fmt.Println("len_s= ", len(s))
	fmt.Println("cap_s= ", cap(s))
	// s[5] = 5
	fmt.Println("s= ", s)
	fmt.Println("")
}

func createEasySlice() {
	a := [5]int{1, 2, 3, 4, 5}
	// s := a[0:2]
	// s := a[2:]
	// s := a[:4]
	s := a[:]
	fmt.Println("s= ", s)
	fmt.Println("")
}

func addSlice() {
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println(s)

	s0 := []int{1, 2, 3}
	s1 := []int{4, 5, 6}
	s2 := append(s0, s1...)
	fmt.Println(s2)
	fmt.Println("")
}

func checkCapacity() {
	s := make([]int, 0, 0)
	fmt.Printf("(A) len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, 1)
	fmt.Printf("(B) len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, []int{2, 3, 4}...)
	fmt.Printf("(C) len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, 5)
	fmt.Printf("(D) len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, 6, 7, 8, 9)
	fmt.Printf("(E) len=%d, cap=%d\n", len(s), cap(s))
}

func main() {
	// createSlice()
	// checkSliceCapa()
	// createEasySlice()
	// addSlice()
	checkCapacity()
}
