package main

import (
	"fmt"
	"sort"
)

func deleteAndEarn(nums []int) int {
	sort.Ints(nums)
	tmpList := []int{nums[0]}
	sum := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			tmpList[len(tmpList)-1] += nums[i]
		} else if nums[i]-1 == nums[i-1] {
			tmpList = append(tmpList, nums[i])
		} else {
			sum += rob(tmpList)
			tmpList = []int{nums[i]}
		}
	}
	sum += rob(tmpList)
	return sum
}

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

	pre := nums[0]
	cur := twoMax(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		pre, cur = cur, twoMax(pre+nums[i], cur)
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
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
}
