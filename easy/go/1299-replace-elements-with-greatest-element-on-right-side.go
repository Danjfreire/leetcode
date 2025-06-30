package main

func replaceElements(arr []int) []int {
	greatest := -1
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], greatest = greatest, max(greatest, arr[i])
	}

	arr[len(arr)-1] = -1

	return arr
}
