package main

import "fmt"

func singleNumber(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	for num, count := range numMap {
		if count == 1 {
			return num
		}
	}
	return -1
}

func main() {
	singleNum := singleNumber([]int{1, 5, 2, 2, 1})
	fmt.Println("出现一次的数： ", singleNum) // Output: 3
}
