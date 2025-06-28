package main

func validMountainArray(arr []int) bool {
	length := len(arr)

	if length < 3 {
		return false
	}

	last := -1
	ascending := true
	maxIdx := 0

	for i, v := range arr {
		if v == last {
			return false
		}

		if ascending {
			if v < last {
				ascending = false
				maxIdx = i - 1
			}
		} else {
			if v > last {
				return false
			}
		}

		last = v
	}

	if maxIdx == 0 || maxIdx == len(arr)-1 {
		return false
	}

	return true
}
