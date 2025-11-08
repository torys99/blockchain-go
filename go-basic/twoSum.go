/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，
并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
*/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("数组下标: %v\n", result) // Output: [0, 1]
	nums = []int{3, 2, 4}
	target = 6
	result = twoSum(nums, target)
	fmt.Printf("数组下标: %v\n", result)
	nums = []int{3, 3}
	target = 6
	result = twoSum(nums, target)
	fmt.Printf("数组下标: %v\n", result)
}
