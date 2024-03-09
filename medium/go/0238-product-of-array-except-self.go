package main

// O(n) without division and with O(1) extra space
func productExceptSelf(nums []int) []int {
    result := make([]int, len(nums), len(nums))
	leftProduct :=1

	for i := 0; i < len(nums); i++ {
		result[i] = leftProduct
		leftProduct *= nums[i]
	}

	rightProduct := 1

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}