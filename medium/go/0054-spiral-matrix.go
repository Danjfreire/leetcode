package main

func spiralOrder(matrix [][]int) []int {

	directions := []string{"right", "down", "left", "up"}

	currentDir := directions[0]

	i := 0
	j := 0
	marker := 1000

	rows := len(matrix) - 1
	cols := len(matrix[0]) - 1

	totalLength := len(matrix) * len(matrix[0])
	res := make([]int, totalLength)
	idx := -1

	for {
		idx++

		if idx == totalLength {
			break
		}

		res[idx] = matrix[i][j]

		if currentDir == "right" {

			reachedBorder := j >= cols || matrix[i][j+1] == marker

			if reachedBorder {
				currentDir = "down"
				i++
			} else {
				matrix[i][j] = marker
				j++
			}

			continue
		}

		if currentDir == "down" {
			reachedBorder := i >= rows || matrix[i+1][j] == marker

			if reachedBorder {
				currentDir = "left"
				j--
			} else {
				matrix[i][j] = marker
				i++
			}

			continue
		}

		if currentDir == "left" {
			reachedBorder := j <= 0 || matrix[i][j-1] == marker

			if reachedBorder {
				currentDir = "up"
				i--
			} else {
				matrix[i][j] = marker
				j--
			}

			continue
		}

		if currentDir == "up" {
			reachedBorder := i <= 0 || matrix[i-1][j] == marker

			if reachedBorder {
				currentDir = "right"
				j++
			} else {
				matrix[i][j] = marker
				i--
			}

			continue
		}
	}

	return res
}

func spiralOrder2(matrix [][]int) []int {

	rows, cols := len(matrix), len(matrix[0])
	res := make([]int, 0, rows*cols)
	top := 0
	bottom := rows - 1
	left := 0
	right := cols - 1

	for len(res) < rows*cols {
		// left to right
		for j := left; j <= right && len(res) < rows*cols; j++ {
			res = append(res, matrix[top][j])
		}
		top++
		// top to bottom
		for i := top; i <= bottom && len(res) < rows*cols; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// right to left
		for j := right; j >= left && len(res) < rows*cols; j-- {
			res = append(res, matrix[bottom][j])
		}
		bottom--
		// bottom to top
		for i := bottom; i >= top && len(res) < rows*cols; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}

func main() {
	matrix := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}

	spiralOrder(matrix)
}
