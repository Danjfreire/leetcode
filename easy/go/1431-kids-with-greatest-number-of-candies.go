// O(n)
func kidsWithCandies(candies []int, extraCandies int) []bool {
    result := make([]bool, len(candies), len(candies));
	maxCandies := 0
	
	for _, c := range candies {
		if c > maxCandies {
			maxCandies = c
		}
	}

	for idx, c := range candies {
		if c + extraCandies >= maxCandies {
			result[idx] = true
		}
	}

	return result
}