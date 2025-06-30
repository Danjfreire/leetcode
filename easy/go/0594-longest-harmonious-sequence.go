package main

import "sort"

// approach 1 - using a map
func findLHS(nums []int) int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	result := 0
	for num, count := range freq {
		if count2, ok := freq[num+1]; ok {
			result = max(result, count+count2)
		}
	}

	return result
}

// approach 2 - two pointer / sliding window
func findLHS2(nums []int) int {
	sort.Ints(nums)
	l, r := 0, 0
	max, currLen := 0, 0

	for r < len(nums) {
		diff := nums[r] - nums[l]
		if diff == 1 {
			currLen = r - l + 1

			if currLen > max {
				max = currLen
			}
		}

		if diff <= 1 {
			r++
		} else {
			l++
		}
	}

	return max
}
