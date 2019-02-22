package main

import (
	"fmt"
)

func recoverFunc() {
	defer func() {
		if x := recover(); x != nil {
			// 変数xはpanicに渡されたinterface
			fmt.Println(x)
		}
	}()

	panic("panic!")
	//これは実行されない
	fmt.Println("Hello, World!")
}

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			// panicによるinterface{}型の値に応じて処理を分岐
			switch v := x.(type) {
			case int:
				fmt.Printf("panic: int=%v\n", v)
			case string:
				fmt.Printf("panic: string=%v\n", v)
			default:
				fmt.Printf("panic: unknown")
			}
		}
	}()
	panic(src)
	return
}

func main() {
	recoverFunc()
	testRecover(128)
	testRecover("hogehoge")
	testRecover([...]int{1, 2, 3})
}
