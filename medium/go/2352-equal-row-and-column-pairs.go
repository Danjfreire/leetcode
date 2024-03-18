package main

import (
	"strconv"
	"strings"
)

func equalPairs(grid [][]int)int {
	rowMap := make(map[string]int)
	n := len(grid)

	for i := 0; i < n; i++ {
		var rowStr  strings.Builder
		for j := 0; j < n; j++ {
			rowStr.WriteString("|")
			rowStr.WriteString(strconv.Itoa(grid[i][j]))
			rowStr.WriteString("|")
		}
		occur, ok := rowMap[rowStr.String()]
		if ok {
			rowMap[rowStr.String()] = occur+1
		} else {
			rowMap[rowStr.String()] = 1
		}
	}

	equalPairs := 0

	for i := 0; i < n; i++ {
		var colStr strings.Builder
		for j := 0; j < n; j++ {
			colStr.WriteString("|")
			colStr.WriteString(strconv.Itoa(grid[j][i]))
			colStr.WriteString("|")
		}

		occur, ok := rowMap[colStr.String()]

		if ok {
			equalPairs += occur
		}
	}

	return equalPairs 
}