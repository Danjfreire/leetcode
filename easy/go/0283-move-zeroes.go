func moveZeroes(nums []int) {
	lastNonZeroIdx := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			temp := nums[lastNonZeroIdx]
			nums[lastNonZeroIdx] = nums[i]
			nums[i] = temp 
			lastNonZeroIdx++;
		}
	}
}

func moveZeroes(nums []int) {
	lastNonZeroIdx := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroIdx], nums[i] = nums[i], nums[lastNonZeroIdx]
			lastNonZeroIdx++
		}
	}
}