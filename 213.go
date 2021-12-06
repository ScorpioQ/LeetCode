package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return twoMax(nums[0], nums[1])
	}

	a := robList(nums[:len(nums)-1])
	b := robList(nums[1:])
	return twoMax(a, b)
}

func robList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return twoMax(nums[0], nums[1])
	}

	pre := nums[0]
	cur := twoMax(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		preX := pre + nums[i]
		pre = cur
		if cur < preX {
			cur = preX
		}
	}

	return cur
}

func twoMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{0}))
}
