package main

import "fmt"

func sum10p(a *int) {
	*a += 10
}

func sum10(a int) {
	a += 10
}

func main() {
	num := 5
	sum10(num)
	fmt.Println("num =", num) // Output: num = 5
	sum10p(&num)
	fmt.Println("num =", num) // Output: num = 15
}
