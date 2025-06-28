package main

import "sort"

type node struct {
	val int
	idx int
}

func maxSubsequence(nums []int, k int) []int {
	heap := make([]node, 0)

	for i, v := range nums {
		heap = append(heap, node{v, i})
	}

	sort.Slice(heap, func(i, j int) bool {
		return heap[i].val < heap[j].val
	})

	// get only the top k elements
	heap = heap[len(nums)-k:]

	// reorder the elements by their original indexes
	sort.Slice(heap, func(i, j int) bool {
		return heap[i].idx < heap[j].idx
	})

	res := make([]int, k)

	for i, v := range heap {
		res[i] = v.val
	}

	return res
}
