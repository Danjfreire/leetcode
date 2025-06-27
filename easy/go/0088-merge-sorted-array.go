package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	last := m + n - 1

	// start at the end of each array, nums1 actually "ends" at m
	p1 := m - 1
	p2 := n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[last] = nums1[p1]
			p1--
		} else {
			nums1[last] = nums2[p2]
			p2--
		}

		last--
	}

	// add any leftover elements from num2 to the start of the list
	for p2 >= 0 {
		nums1[last] = nums2[p2]
		p2--
		last--
	}
}
