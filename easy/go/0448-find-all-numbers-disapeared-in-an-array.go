package main

func findDisappearedNumbers(nums []int) []int {
	appearances := make(map[int]bool)

	count := 0
	for i := 0; i < len(nums); i++ {
		appearances[nums[i]] = true
		count++
	}

	res := make([]int, len(nums)-count)

	for i := 1; i <= len(nums); i++ {
		if _, ok := appearances[i]; !ok {
			res = append(res, i)
		}
	}

	return res
}

// Faster
func findDisappearedNumbers2(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	res := make([]int, length)

	// using result array as sorted array for seen values
	for _, v := range nums {
		res[v-1] = v
	}

	// here we have res array where missing records will be with value = 0
	// scan all values with 0 and "shift" them to head of res array
	j := 0
	for i := 0; i < length; i++ {
		if res[i] == 0 {
			res[j] = i + 1
			j++
		}
	}

	// chop head of res array and return it
	return res[:j]
}
