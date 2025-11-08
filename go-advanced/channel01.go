/*
编写一个程序，使用通道实现两个协程之间的通信。
一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
*/
package main

import (
	"fmt"
)

func generateNums(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
}

func receiveAndPrintNums(ch <-chan int) {
	for i := 1; i <= 100; i++ {
		num := <-ch
		fmt.Println("Received number:", num)
	}
}
func main() {
	ch := make(chan int, 10)
	go generateNums(ch)
	go receiveAndPrintNums(ch)
	var input string
	fmt.Scanln(&input)
}
