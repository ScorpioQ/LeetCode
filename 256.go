package main

import "fmt"

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minCost(costs [][]int) int {
	for i := len(costs) - 2; i >= 0; i-- {
		costs[i][0] += minInt(costs[i+1][1], costs[i+1][2])
		costs[i][1] += minInt(costs[i+1][0], costs[i+1][2])
		costs[i][2] += minInt(costs[i+1][0], costs[i+1][1])
	}
	return minInt(costs[0][0], minInt(costs[0][1], costs[0][2]))
}

func main() {
	fmt.Println(minCost([][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 19},
	}) == 10)
}
