package main

func largestAltitude(gain []int) int {
	currentAltitude := 0
	maxAltitude := 0 

	for i := 0 ; i < len(gain); i++ {
		currentAltitude += gain[i]
		if currentAltitude > maxAltitude {
			maxAltitude = currentAltitude
		}
	}

	return maxAltitude
}