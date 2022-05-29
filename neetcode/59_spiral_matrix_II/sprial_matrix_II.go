package spiral_matrix_II

func generateMatrix(n int) [][]int {
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, n)
	}

	counter := 1

	left, right := 0, len(nums[0])-1
	top, bottom := 0, len(nums)-1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			nums[top][i] = counter
			counter++
		}
		top++

		for i := top; i <= bottom; i++ {
			nums[i][right] = counter
			counter++
		}
		right--

		if left > right || top > bottom {
			break
		}

		for i := right; i >= left; i-- {
			nums[bottom][i] = counter
			counter++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			nums[i][left] = counter
			counter++
		}
		left++
	}

	return nums
}
