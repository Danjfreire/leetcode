package main

func removeDuplicates(nums []int) int {
	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
