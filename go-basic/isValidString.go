/*
*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
package main

import "fmt"

func isValidString(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
func main() {
	testStr := "{[()]}"
	result := isValidString(testStr)
	fmt.Printf("字符串 %s 是否有效: %v\n", testStr, result) // Output: true
	testStr = "{[(])}"
	result = isValidString(testStr)
	fmt.Printf("字符串 %s 是否有效: %v\n", testStr, result) // Output: false
	testStr = "{{[[(())]]}}"
	result = isValidString(testStr)
	fmt.Printf("字符串 %s 是否有效: %v\n", testStr, result) // Output: true
}
