package main

import (
	"fmt"
)

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/
func isPalindrome(x int) bool {
	if x < 0 { // 负数不是回文数
		return false
	}
	str := fmt.Sprintf("%d", x)
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	num := 121
	result := isPalindrome(num)
	fmt.Printf("数字 %d 是否是回文数: %v\n", num, result) // Output: true
	num = -121

	result = isPalindrome(num)
	fmt.Printf("数字 %d 是否是回文数: %v\n", num, result) // Output: false
	num = 10
	result = isPalindrome(num)
	fmt.Printf("数字 %d 是否是回文数: %v\n", num, result) // Output: false
}
