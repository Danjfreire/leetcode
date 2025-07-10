package main

func dominantIndex(nums []int) int {
	largest := -1
	largestIdx := -1
	secondLargest := -1

	for k, v := range nums {
		if v > largest {
			secondLargest = largest
			largest = v
			largestIdx = k
		} else if v > secondLargest {
			secondLargest = v
		}
	}

	if largest == 0 {
		return -1
	}

	if secondLargest*2 <= largest {
		return largestIdx
	} else {
		return -1
	}
}
