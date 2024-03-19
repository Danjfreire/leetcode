package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
			continue
		}

		str := ""
		for stack[len(stack)-1] != '[' {
			str = string(stack[len(stack)-1]) + str
			stack = stack[:len(stack)-1]
		}

		stack = stack[:len(stack)-1]

		k := ""
		for len(stack) > 0 {
            _, err := strconv.Atoi(string(stack[len(stack)-1]))

			if err != nil {
				break;
			}

			k = string(stack[len(stack)-1]) + k
			stack = stack[:len(stack)-1]
		}

		if num, err := strconv.Atoi(k); err == nil {
         parsed := strings.Repeat(str, num)
		 for i := 0; i < len(parsed); i++ {
			stack = append(stack, parsed[i])
		 }
		} else {
			return err.Error();
		}
	}

	return string(stack)
}

