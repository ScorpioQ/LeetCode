package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 3 {
		return n
	}

	pre := 2
	cur := 3
	for i := 4; i <= n; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}
