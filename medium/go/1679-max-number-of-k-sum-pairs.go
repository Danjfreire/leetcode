package main

// O(n)
func maxOperations(nums []int, k int) int {
	numMap := make(map[int]int, len(nums))
	operations := 0

	for _, v := range nums {
		if v >= k {
			continue
		}

		dif := k - v

		if numMap[dif] != 0 {
			numMap[dif]--
			operations++
			continue
		}
		numMap[v]++
    }
 
	return operations 
}