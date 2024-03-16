package main

import (
	"slices"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	letters := make(map[byte]bool)
	word1LetterMap := make(map[byte]int)
	word2LetterMap := make(map[byte]int)

	for i := 0; i < len(word1); i++ {
		word1Occur,ok1 := word1LetterMap[word1[i]];
		word2Occur,ok2 := word2LetterMap[word2[i]];

		if ok1 {
			word1LetterMap[word1[i]] = word1Occur + 1;
		} else {
			word1LetterMap[word1[i]] = 1;
		}

		if ok2 {
			word2LetterMap[word2[i]] = word2Occur + 1;
		} else {
			word2LetterMap[word2[i]] = 1;
		}

		letters[word1[i]] = true
		letters[word2[i]] = true
	}

    // check if the words have the same letters
	for letter := range letters {
		_, ok1 := word1LetterMap[letter]
		_, ok2 := word2LetterMap[letter]

		if !ok1 || !ok2 {
			return false
		}
	}

	// check if the occurrences have the same distribution
	word1Occurrences := getMapValues(word1LetterMap) 
	word2Occurrences := getMapValues(word2LetterMap) 

	slices.Sort(word1Occurrences)
	slices.Sort(word2Occurrences)

	for i := range word1Occurrences {
		if word1Occurrences[i] != word2Occurrences[i] {
			return false
		}
	}

	return true
}

func getMapValues(m map[byte]int) []int {
	occurrences := []int{}

	for _, v := range m {
		occurrences = append(occurrences, v)
	}

	return occurrences
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    word1Occurrences := make([]int, 26)
    word2Occurrences := make([]int, 26)

    for i:=0; i<len(word1); i++ {
        char1 := word1[i]
        charCode1 := int('z'-char1)
        word1Occurrences[charCode1] = word1Occurrences[charCode1] + 1

        char2 := word2[i]
        charCode2 := int('z'-char2)
        word2Occurrences[charCode2] = word2Occurrences[charCode2] + 1
    }

	// check if the letters are the same
    for i:=0; i<len(word1Occurrences); i++ {
        if !((word1Occurrences[i] > 0 && word2Occurrences[i]> 0) || 
		(word1Occurrences[i] == 0 && word2Occurrences[i] == 0)) {
            return false
        }
    }

	// check if the occurrences have the same distribution
    slices.Sort(word1Occurrences)
    slices.Sort(word2Occurrences)

    for i:=0; i<len(word1Occurrences); i++ {
        if word1Occurrences[i] != word2Occurrences[i] {
            return false
        }
    }

    return true
}