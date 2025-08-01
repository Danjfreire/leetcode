package main

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {
	input := []int{0, 1, 2, 2, 3, 0, 4, 2}

	res := removeElement(input, 2)
	fmt.Println(input)
	fmt.Println(res)
}
