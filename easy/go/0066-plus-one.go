package main

import "fmt"

func plusOne(digits []int) []int {
	carry := 1
	i := len(digits) - 1

	for carry != 0 && i >= 0 {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			carry = 0
		}

		i--
	}

	if carry == 1 {
		res := make([]int, len(digits)+1)
		copy(res[1:], digits)
		res[0] = 1
		return res
	}

	return digits
}

func main() {
	res := plusOne([]int{9})

	fmt.Println(res)
}
