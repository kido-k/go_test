package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("My", "name", "is", "Taro")
	fmt.Print("\n") //改行

	fmt.Print("Hello, World!")
	fmt.Print("My", "name", "is", "Taro", "\n")
	fmt.Print("\n") //改行

	// 10進数の形式で数値5を%dの箇所へ埋め込む
	fmt.Printf("数値=%d\n", 5) // => "数値=5"
	fmt.Print("\n")          //改行

	// 数値用の書式いろいろ
	fmt.Printf("10進数=%d 2進数=%b\n 8進数=%o 16進数=%x\n", 17, 17, 17, 17)
	// => "10進数=17 2進数=10001 8進数=21 16進数=11"
	fmt.Print("\n") //改行

	// 埋め込むパラメータが足りない
	fmt.Printf("%d年%d月%d日\n", 2019, 2)
	// => "10進数=17 2進数=10001 8進数=21 16進数=11"
	fmt.Print("\n") //改行

	// %vはさまざまな型のデータを埋め込む
	fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3})
	// => "数値=5 文字列=Golang 配列=[1,2,3]"
	fmt.Print("\n") //改行

	// %#vはGoのリテラル表現でデータを埋め込む
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3})
	// => "数値=5 文字列="Golang" 配列=[3][1,2,3]"
	fmt.Print("\n") //改行

	// %Tはデータの型情報を埋め込む
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [...]int{1, 2, 3})
	// => "数値=int 文字列=%T 配列=[3]int"
	fmt.Print("\n") //改行
}
