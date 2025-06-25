package main

// Obvious solution using language features
// func sortedSquares(nums []int) []int {
// 	for i, v := range nums {
// 		nums[i] = v * v
// 	}

// 	slices.Sort(nums)

// 	return nums
// }

// solution using 2 pointers (faster)

func Abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	left, right := 0, len(nums)-1

	for i := len(nums) - 1; i >= 0; i-- {
		var num int

		absR := Abs(nums[right])
		absL := Abs(nums[left])

		if absR > absL {
			num = absR
			right--
		} else {
			num = absL
			left++
		}

		res[i] = num * num
	}

	return res
}
