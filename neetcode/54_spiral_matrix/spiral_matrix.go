package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	nums := []int{}

	left, right := 0, len(matrix[0])-1
	top, bottom := 0, len(matrix)-1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			nums = append(nums, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			nums = append(nums, matrix[i][right])
		}
		right--

		if left > right || top > bottom {
			break
		}

		for i := right; i >= left; i-- {
			nums = append(nums, matrix[bottom][i])
		}
		bottom--

		for i := bottom; i >= top; i-- {
			nums = append(nums, matrix[i][left])
		}
		left++
	}

	return nums
}
