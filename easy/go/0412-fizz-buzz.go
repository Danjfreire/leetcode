package main

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string, n)

	for i := 0; i < n; i++ {
		num := i + 1

		if num%5 == 0 && num%3 == 0 {
			res[i] = "FizzBuzz"
		} else if num%3 == 0 {
			res[i] = "Fizz"
		} else if num%5 == 0 {
			res[i] = "Buzz"
		} else {
			res[i] = strconv.Itoa(num)
		}
	}

	return res
}
