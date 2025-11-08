package main

import "fmt"

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func multiply(a *[]int) {
	for i, v := range *a {
		(*a)[i] = v * 2
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Before multiply:", nums)
	multiply(&nums)
	fmt.Println("After multiply:", nums)
}
