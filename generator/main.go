package main

import (
	f "fmt"
)

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	ints := integers()
	f.Println(ints())
	f.Println(ints())
	f.Println(ints())

	other_ints := integers()
	f.Println(other_ints())
}
