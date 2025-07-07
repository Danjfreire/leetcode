package main

import (
	"slices"
)

func heightChecker(heights []int) int {
	cpy := make([]int, len(heights))

	copy(cpy, heights)

	slices.SortFunc(cpy, func(i, j int) int {
		if i <= j {
			return -1
		}

		return 1
	})

	res := 0

	for i := 0; i < len(heights); i++ {
		if heights[i] != cpy[i] {
			res++
		}
	}

	return res
}
