package main

import (
	"fmt"
	"runtime"
)

func sub() {
	// for {
	// 	fmt.Println("sub loop")
	// }

	for i := 0; i < 10; i += 1 {
		fmt.Println("sub loop")
	}
}

func main() {
	go sub()
	// for {
	// 	fmt.Println("main loop")
	// }

	fmt.Println("Year") //追加でGoroutineが増える
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())
}
