package main

import (
	"math"
)

func findKDistantIndices(nums []int, key int, k int) []int {
	keyIndexes := []int{}

	for i, v := range nums {
		if v == key {
			keyIndexes = append(keyIndexes, i)
		}
	}

	res := []int{}

	if len(keyIndexes) == 0 {
		return res
	}

	for i, _ := range nums {
		for _, j := range keyIndexes {
			if math.Abs(float64(i-j)) <= float64(k) {
				res = append(res, i)
				break
			}
		}
	}

	return res
}
