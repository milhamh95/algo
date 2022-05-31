package search_a_2d_matrix

// brute force => O(n * n => n ^ 2)
func searchMatrix(matrix [][]int, target int) bool {
	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[0]); column++ {
			num := matrix[row][column]

			if num == target {
				return true
			}

		}
	}

	return false
}

// O(m + n)
// two pointer approach
func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	startColumn := 0
	endColumn := len(matrix[0]) - 1

	startRow := 0
	endRow := len(matrix) - 1

	for startRow <= endRow && endColumn >= startColumn {
		num := matrix[startRow][endColumn]

		if target == num {
			return true
		}

		if target > num {
			startRow++
			continue
		}

		endColumn--
	}

	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	startColumn := 0
	endColumn := len(matrix)*len(matrix[0]) - 1

	totalColumn := len(matrix[0]) - 1

	for startColumn <= endColumn {
		mid := startColumn + (endColumn-startColumn)/2
		row := mid / totalColumn
		column := mid % totalColumn

		num := matrix[row][column]

		if target == num {
			return true
		}

		if target > num {
			startColumn = mid + 1
			continue
		}

		if target < num {
			endColumn = mid - 1
		}
	}

	return false
}
