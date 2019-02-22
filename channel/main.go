package main

import (
	"fmt"
	"time"
)

// func checkChannel() {

// 	var (
// 		ch0 chan int
// 		ch1 <-chan int
// 		ch2 chan<- int
// 	)
// 	ch1 = ch0
// 	ch2 = ch0
// 	ch0 = ch1
// 	ch0 = ch2
// 	ch1 = ch2
// 	ch2 = ch1

// }

func receive(name string, ch <-chan int) {

	for {
		i, ok := <-ch
		if ok == false {
			// 受信できなくなったら終了
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func checkGoroutine() {
	ch := make(chan int, 20)

	go receive("1st goroutine", ch)
	go receive("2nd goroutine", ch)
	go receive("3rd goroutine", ch)

	i := 0
	for i < 200 {
		ch <- i
		i += 1
	}
	close(ch)

	// ゴルーチンの完了を3待つ
	time.Sleep(3 * time.Second)
}

func checkSelect() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2

	select {
	case <-ch1:
		fmt.Println("ch1から受信")
	case <-ch2:
		fmt.Println("ch1から受信")
	case <-ch3:
		fmt.Println("ch3から受信")
	default:
		fmt.Println("ここへは到達しない")
	}
}

func checkSelectExample() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)
		}
	}()

	go func() {
		for {
			i := <-ch2
			ch3 <- i - 1
		}
	}()

	n := 1
LOOP:
	for {
		select {
		case ch1 <- n:
			n++
		case i := <-ch3:
			fmt.Println("received", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}
}

func main() {
	// checkChannel()
	// checkGoroutine()
	// checkSelect()
	checkSelectExample()
}
