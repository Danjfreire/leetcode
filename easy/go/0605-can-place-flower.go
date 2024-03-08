// O(n)
func canPlaceFlowers(flowerbed []int, n int) bool {
	plantedFlowers := 0

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}

		hasLeft := i > 0 && flowerbed[i-1] == 1
		hasRight := i < len(flowerbed)-1 && flowerbed[i+1] == 1

		if !hasLeft && !hasRight {
			flowerbed[i] = 1
			plantedFlowers++

            if plantedFlowers >= n {
                return true
            }
		}
	}

	return plantedFlowers >= n
}