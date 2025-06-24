package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	magazineMap := make(map[rune]int)

	for _, v := range magazine {
		magazineMap[v]++
	}

	for _, v := range ransomNote {
		val, ok := magazineMap[v]

		if !ok || val == 0 {
			return false
		}

		magazineMap[v]--
	}

	return true
}
