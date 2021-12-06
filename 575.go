package main

import "fmt"

func distributeCandies(candyType []int) int {
	save := map[int]int{}
	for _, item := range candyType {
		save[item] += 1
	}
	if len(candyType)/2 < len(save) {
		return len(candyType) / 2
	} else {
		return len(save)
	}
}

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
}
