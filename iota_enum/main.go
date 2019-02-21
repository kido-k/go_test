package main

import (
	f "fmt"
)

func checkIota1() {
	f.Println("check iota 1")

	const A = iota
	const B = iota
	const C = iota

	const D = 1 + iota
	const E = 1 + iota
	const F = 1 + iota

	const G = 17
	const H = 17 + iota
	const R = iota

	f.Println(A)
	f.Println(B)
	f.Println(C)

	f.Println(D)
	f.Println(E)
	f.Println(F)

	f.Println(G)
	f.Println(H)
	f.Println(R)
}

func checkIota2() {
	f.Println("check iota 2")

	const (
		A = iota
		B = iota
		C = iota

		D = 1 + iota
		E = 1 + iota
		F = 1 + iota

		G = 17
		H = 17 + iota
		R = iota
	)

	f.Println(A)
	f.Println(B)
	f.Println(C)

	f.Println(D)
	f.Println(E)
	f.Println(F)

	f.Println(G)
	f.Println(H)
	f.Println(R)
}

func main() {
	checkIota1()
	checkIota2()
}
