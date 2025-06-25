package main

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	streak := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			streak++
		} else {
			streak = 0
		}

		if streak > max {
			max = streak
		}
	}

	return max
}
