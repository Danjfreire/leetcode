package main

func runningSum(nums []int) []int {
	sum := 0
	res := make([]int, len(nums))

	for i, v := range nums {
		sum += v
		res[i] = sum
	}

	return res
}
