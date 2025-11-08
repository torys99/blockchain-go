package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs[1:] {
		if len(prefix) > len(str) {
			prefix = prefix[:len(str)]
		}
		for len(prefix) > 0 && str[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Printf("最长公共前缀: %s\n", result) // Output: "fl"
	strs = []string{"dog", "racecar", "car"}
	result = longestCommonPrefix(strs)
	fmt.Printf("最长公共前缀: %s\n", result) // Output: ""
}
