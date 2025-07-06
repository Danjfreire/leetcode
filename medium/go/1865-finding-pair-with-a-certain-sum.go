package main

type FindSumPairs struct {
	nums1    []int
	nums2    []int
	nums2Map map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	nums2Map := make(map[int]int)
	for _, v := range nums2 {
		nums2Map[v]++
	}

	pairs := FindSumPairs{nums1, nums2, nums2Map}

	return pairs
}

func (this *FindSumPairs) Add(index int, val int) {
	prev := this.nums2[index]
	this.nums2Map[prev]--
	this.nums2[index] += val
	this.nums2Map[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	count := 0

	for _, v := range this.nums1 {
		diff := tot - v
		if val, ok := this.nums2Map[diff]; ok {
			count += val
		}
	}

	return count
}
