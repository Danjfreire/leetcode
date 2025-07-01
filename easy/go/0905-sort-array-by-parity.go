package main

func sortArrayByParity(nums []int) []int {
	write := 0

	for read := 0; read < len(nums); read++ {
		writeIsEven := nums[write]%2 == 0
		readIsEven := nums[read]%2 == 0

		// if there is a change to be made
		if !writeIsEven && readIsEven {
			nums[write], nums[read] = nums[read], nums[write]
			write++
		} else if writeIsEven && readIsEven { // if there is no change to be made
			write++
		}

		// if can't change because both are odd, don't increase write
	}

	return nums
}
