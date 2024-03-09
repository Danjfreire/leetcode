package main

import "strconv"

// O(n) time and O(1) extra space
func compress(chars []byte) int {
    idx := 0
	resultIdx := 0

	for idx < len(chars) {
		char := chars[idx]
		sequenceLen := 1

		for idx+sequenceLen < len(chars) && chars[idx + sequenceLen] == char {
			sequenceLen++
		}

		idx += sequenceLen

		chars[resultIdx] = char
		if sequenceLen > 1 {
			resultIdx++
			lenStr := strconv.Itoa(sequenceLen)
			for i := 0; i < len(lenStr); i ++ {
				chars[resultIdx] = lenStr[i]
				resultIdx++
			}
		} else {
			resultIdx++
		}
	}

	return resultIdx
}