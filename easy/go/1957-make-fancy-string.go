package main

import (
	"fmt"
	"strings"
)

func makeFancyString(s string) string {
	res := strings.Builder{}

	res.WriteByte(s[0])
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			count = 1
		} else {
			count++
		}

		if count < 3 {
			res.WriteByte(s[i])
		}
	}

	return res.String()
}

func main() {
	fmt.Println(makeFancyString("leeetcode"))
	fmt.Println(makeFancyString("aaabaaaa"))
	fmt.Println(makeFancyString("aab"))
}
