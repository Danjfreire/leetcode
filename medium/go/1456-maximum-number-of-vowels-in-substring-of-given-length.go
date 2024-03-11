package main

func maxVowels(s string, k int) int {
	vowels := map[byte]bool{
		'a':true,
		'e':true,
		'i':true,
		'o':true,
		'u':true,
	}

	currentVowels := 0 

	for i := 0; i < k; i++ {
		_, has := vowels[s[i]]

		if has {
			currentVowels++
		}
	}

	max := currentVowels

	for i:= k; i < len(s) ; i++ {
		if _, has := vowels[s[i]]; has {
			currentVowels++
		}

		if _, has := vowels[s[i-k]]; has {
			currentVowels--
		}

		if currentVowels > max {
			max = currentVowels
		}
	}

	return max 
}