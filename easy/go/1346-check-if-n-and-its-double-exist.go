package main

func checkIfExist(arr []int) bool {
	table := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		table[arr[i]] = i
	}

	for i := 0; i < len(arr); i++ {
		double := arr[i] * 2
		idx, ok := table[double]

		if ok && idx != i {
			return true
		}
	}

	return false
}
