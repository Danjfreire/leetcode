package main

func duplicateZeros(arr []int) {
	i := 0
	for i < len(arr) {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}

			if i+1 < len(arr) {
				arr[i+1] = 0
			}
			i += 2
		} else {
			i++
		}
	}
}
