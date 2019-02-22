package main

import (
	"fmt"
)

func main() {
	flag_a := true
	flag_b := false
	flag_c := true
	num_a := 1
	num_b := 0
	// flag_c := true
	// flag_d := true

	if flag_a {
		fmt.Println("flag_a")
	} else if flag_c {
		fmt.Println("flag_c")
	} else {
		fmt.Println("flag_b")
	}

	if flag_b {
		fmt.Println("flag_b")
	}

	// if num_a {
	if num_a == 1 {
		fmt.Println("num_a")
	}

	// if num_b {
	if num_b == 1 {
		fmt.Println("num_b")
	}
}
