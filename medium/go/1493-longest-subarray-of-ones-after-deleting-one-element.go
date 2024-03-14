package main

func longestSubarray(nums []int) int {
	left := 0
	right := 0
	lastZeroIdx := -1
	maxLen := 0

	for right < len(nums) {
		if nums[right] == 0 && lastZeroIdx == -1 {
			lastZeroIdx = right
		} else if (nums[right] == 0) {
			maxLen = max(maxLen, right-1-left)
			left = lastZeroIdx+1
			lastZeroIdx = right
		} else {
			maxLen = max(maxLen, right-left)
		}

		right++
	}

	return maxLen
}

func max(v1 int, v2 int) int {
	if v1>v2 {
		return v1
	}
	return v2
}