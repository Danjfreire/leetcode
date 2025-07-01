package main

func possibleStringCount(word string) int {
	count := 1

	prev := 0
	for i := 1; i < len(word); i++ {
		if word[i] == word[prev] {
			count++
		}
		prev++
	}

	return count
}
