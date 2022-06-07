package daily_temperature

// 1st => brute force => O(n ^ 2)
// 2nd => O(n)
// because inner loop will never run more than N times total
// if usual quadratic, then inner loop call more than N times total
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)

	// store number days have to wait to get a warmer temperature
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 0
	}

	stack := []int{}
	for i, v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			idxPrevNumber := stack[len(stack)-1]
			numDaysToWait := i - idxPrevNumber
			result[idxPrevNumber] = numDaysToWait

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return result
}
