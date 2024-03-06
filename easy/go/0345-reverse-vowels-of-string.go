// O(n)
func reverseVowels(s string) string {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	strVowels := []byte{}

	for i:= 0; i < len(s); i++ {
		if slices.Contains(vowels, s[i]) {
			strVowels = append(strVowels, s[i]);
		}
	}
    
	var result strings.Builder
	vowelIdx := len(strVowels) - 1;

	for i := 0; i < len(s); i++ {
		if slices.Contains(vowels, s[i]) {
			result.WriteByte(strVowels[vowelIdx])
			strVowels = strVowels[:vowelIdx]
			vowelIdx--
		} else {
			result.WriteByte(s[i])
		}
	}

	return result.String()
}

func reverseVowels(s string) string {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	buffer := []rune(s)
	foundVowels := []rune{}

	for i:= 0; i < len(buffer); i++ {
		if slices.Contains(vowels, buffer[i]) {
			foundVowels = append(foundVowels, buffer[i]);
		}
	}

	vowelIdx := 0

	for i := len(buffer)-1; i >= 0; i-- {
		if slices.Contains(vowels, buffer[i]) { 
			buffer[i] = foundVowels[vowelIdx]
			vowelIdx++
		}
	}

	return string(buffer) 
}