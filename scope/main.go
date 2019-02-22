package main

import (
	"fmt"

	"./foo"
)

func main() {

	fmt.Println(foo.MAX)
	// fmt.Println(foo.internal_max)

	fmt.Println(foo.FooFunc(5))
	// fmt.Println(foo.internalFunc(5))
}
