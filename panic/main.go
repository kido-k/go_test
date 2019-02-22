package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Defer")

	panic("runtime error!") //ここでエラー終了
	fmt.Println("Hello, World!")
}
