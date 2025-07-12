package main

func findDiagonalOrder(mat [][]int) []int {
	i := 0
	j := 0
	rows := len(mat) - 1
	cols := len(mat[0]) - 1
	res := make([]int, len(mat)*len(mat[0]))
	idx := 0
	up := true

	for {
		// terminating condition
		if i+j > rows+cols {
			break
		}

		res[idx] = mat[i][j]

		if up && j >= cols { // right border
			up = false
			i++
		} else if up && i <= 0 { // top border
			up = false
			j++
		} else if !up && i >= rows { // bottom border
			up = true
			j++
		} else if !up && j <= 0 { // left border
			up = true
			i++
		} else if up { // regular case going up
			i--
			j++
		} else { // regular case going down
			i++
			j--
		}

		idx++
	}
	return res
}
