package main

import "fmt"

func twoSum(nums []int, target int) []int {
	save := map[int]int{}
	for idx, item := range nums {
		if val, ok := save[item]; ok {
			return []int{val, idx}
		}
		save[target-item] = idx
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
