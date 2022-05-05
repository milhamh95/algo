package matrix

func Matrix(n int) [][]int {
	// initalize 2d slice with row = n
	result := make([][]int, n)
	counter := 1

	// insert column on each row
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	startRow := 0
	endRow := n - 1

	startColumn := 0
	endColumn := n - 1

	for startColumn <= endColumn &&
		startRow <= endRow {
		for i := startColumn; i <= endColumn; i++ {
			result[startRow][i] = counter
			counter++
		}
		startRow++

		for i := startRow; i <= endRow; i++ {
			result[i][endColumn] = counter
			counter++
		}
		endColumn--

		for i := endColumn; i >= startColumn; i-- {
			result[endRow][i] = counter
			counter++
		}
		endRow--

		for i := endRow; i >= startRow; i-- {
			result[i][startColumn] = counter
			counter++
		}
		startColumn++
	}

	return result
}
