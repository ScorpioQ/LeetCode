package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	pre := cost[0]
	cur := cost[1]

	for i, c := range cost {
		if i < 2 {
			continue
		}
		preX := pre + c
		curX := cur + c
		pre = cur
		if preX < curX {
			cur = preX
		} else {
			cur = curX
		}
	}

	if pre < cur {
		cur = pre
	}

	return cur
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
