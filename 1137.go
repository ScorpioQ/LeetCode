package main

import "fmt"

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	prepre := 0
	pre := 1
	cur := 1
	for i := 3; i <= n; i++ {
		prepre, pre, cur = pre, cur, cur+pre+prepre
	}

	return cur
}

func main() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
}
