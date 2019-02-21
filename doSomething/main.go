package main

import (
	f "fmt"
)

func doSomething() (x, y int) {
	y = 5
	return
}

func main() {
	f.Println(doSomething())
}
