package main

import "math"

// O(n) and O(1) space
func increasingTriplet(nums []int) bool {
    min := math.MaxInt
	mid := math.MaxInt

	for _, v := range nums {
		if v > mid {
			return true
		}

		if v > min {
			mid = v
		} else {
			min = v
		}
	}

	return false
}