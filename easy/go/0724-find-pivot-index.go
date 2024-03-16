package main

func pivotIndex(nums []int) int {
	totalSum := 0

	for i := 0; i < len(nums); i++ {
		totalSum+= nums[i]
	}

	leftSum := 0
	for i := 0; i < len(nums); i++ {
	  rightSum := totalSum - nums[i] - leftSum
	  if rightSum == leftSum {
		return i
	  } else {
		leftSum += nums[i]
	  }
	}

	return -1
}