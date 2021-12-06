package main

import "fmt"

func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if max < i {
			return false
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return max >= len(nums)-1
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
