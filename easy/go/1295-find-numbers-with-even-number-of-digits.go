package main

import "strconv"

func findNumbers(nums []int) int {
	count := 0

	for _, v := range nums {
		conv := strconv.Itoa(v)
		if len(conv)%2 == 0 {
			count++
		}
	}

	return count
}
