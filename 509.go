package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	pre := 1
	cur := 1
	for i := 3; i <= n; i++ {
		pre, cur = cur, cur+pre
	}

	return cur
}

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
}
