package main

import "sort"

func thirdMax(nums []int) int {
	sort.Ints(nums)
	max := nums[len(nums)-1]
	maxMap := make(map[int]int)

	for i := len(nums) - 1; i >= 0; i-- {
		maxMap[nums[i]]++

		if len(maxMap) == 3 {
			return nums[i]
		}
	}

	return max
}
