package main

import (
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	reg := regexp.MustCompile("\\s+")
	words := reg.Split(s, -1)
	var result strings.Builder

	for i := len(words) - 1; i >=0; i-- {
		result.WriteString(words[i])
		result.WriteString(" ")
	}

	return strings.TrimSpace(result.String()) 
}