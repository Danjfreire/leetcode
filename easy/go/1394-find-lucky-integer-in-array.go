package main

func findLucky(arr []int) int {
	freq := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		freq[arr[i]]++
	}

	largest := -1
	for k, v := range freq {
		if k == v {
			if k > largest {
				largest = k
			}
		}
	}

	return largest
}
