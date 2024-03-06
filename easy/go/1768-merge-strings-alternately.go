
func Max (x int, y int) int {
	if x > y {
		return x
	}

	return y
}
func mergeAlternately(word1 string, word2 string) string {
	len1 := len(word1)
	len2 := len(word2)
	maxLen := Max(len1, len2) 
	var result strings.Builder 

	for i := 0; i < maxLen; i++ {
		if i < len1 {
			result.WriteByte(word1[i])
		} 

		if i < len2 {
			result.WriteByte(word2[i])
		}
	}

	return result.String()
}