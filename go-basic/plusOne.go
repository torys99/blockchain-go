/*
*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
将大整数加 1，并返回结果的数字数组
*/
package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	newDigits := make([]int, n+1)
	newDigits[0] = 1
	return newDigits
}

func main() {
	digits := []int{1, 2, 3}
	result := plusOne(digits)
	fmt.Printf("加一后的结果: %v\n", result) // Output: [1, 2, 4]
	digits = []int{4, 3, 2, 1}
	result = plusOne(digits)
	fmt.Printf("加一后的结果: %v\n", result) // Output: [1, 0, 0, 0]
	digits = []int{9, 9}
	result = plusOne(digits)
	fmt.Printf("加一后的结果: %v\n", result)
}
