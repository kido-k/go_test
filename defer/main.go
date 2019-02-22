package main

import (
	"fmt"
)

func runDeferString() {
	defer fmt.Println("defer")
	fmt.Println("done")
}

func runDeferNumber() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("done")
}

// func runOpenFile() {
// 	file, err := os.Open("/path/to/file")
// 	if err != nil {
// 		// ファイルのオープンに失敗したらreturn
// 		return
// 	}
// 	// ファイルのクローズ処理を登録
// 	defer file.Close()
// }

func main() {
	runDeferString()
	fmt.Println("")
	runDeferNumber()
	fmt.Println("")
	// runOpenFile()
	// fmt.Println("")
}
