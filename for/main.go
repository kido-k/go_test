package main

import (
	"fmt"
)

func fora() {
	var i int
	i = 0

	fmt.Println("fora")
	for {
		fmt.Println(i)
		i += 1
		if i == 10 {
			break
		}
	}
	fmt.Println("")
}

func forb() {
	fmt.Println("forb")
	for i := 0; i < 10; i += 1 {
		fmt.Println(i)
	}
	fmt.Println("")
}

func forfruits() {
	fruits := [3]string{"Apple", "Banana", "Cherry"}

	fmt.Println("forfruits")
	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}
	fmt.Println("")
}

func main() {
	fora()
	forb()
	forfruits()
}
