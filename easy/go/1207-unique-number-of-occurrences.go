package main

func uniqueOccurrences(arr []int) bool{
	occurrenceMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		occurrences, ok := occurrenceMap[arr[i]]
		if ok {
			occurrenceMap[arr[i]] = occurrences + 1
		} else {
			occurrenceMap[arr[i]] = 1
		}
	}

	occurrenceSet := make(map[int]bool)

	for _, v := range occurrenceMap {
		_, ok := occurrenceSet[v];

		if ok {
			return false
		} else {
			occurrenceSet[v] = true
		}
	}

	return true
}