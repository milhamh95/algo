package trapping_rain_water

// time complexity -> O(n ^ 2)
// space complexity -> O(1)
func trap1(height []int) int {
	n := len(height)
	maxTrappedWater := 0

	if n < 3 {
		return maxTrappedWater
	}

	// loop wall from index 1
	for i := 1; i < n-1; i++ {
		currentWall := height[i]
		maxLeft := currentWall
		maxRight := currentWall

		// search for max left wall, from beginning
		// until wall i
		for j := 0; j < i; j++ {
			currentLeftWall := height[j]
			maxLeft = max(maxLeft, currentLeftWall)
		}

		// search for max right wall, from i + 1
		// until the end
		for j := i + 1; j < n; j++ {
			currentRightWall := height[j]
			maxRight = max(maxRight, currentRightWall)
		}

		trappedWater := min(maxLeft, maxRight) - currentWall
		maxTrappedWater += trappedWater
	}

	return maxTrappedWater
}

// time complexity => O(n)
// space complexity => (n)
func trap2(height []int) int {
	maxTrappedWater := 0
	n := len(height)

	if n < 3 {
		return maxTrappedWater
	}

	maxLeft := make([]int, n)
	maxRight := make([]int, n)

	maxLeft[0] = height[0]
	maxRight[n-1] = height[n-1]

	// search max left wall
	for i := 1; i < n; i++ {
		currentWall := height[i]
		maxLeft[i] = max(currentWall, maxLeft[i-1])
	}

	// search max right wall
	for i := n - 2; i >= 0; i-- {
		currentWall := height[i]
		maxRight[i] = max(currentWall, maxRight[i+1])
	}

	// calculate trapped water each wall
	for i := 0; i < n; i++ {
		trappedWater := min(maxLeft[i], maxRight[i]) - height[i]
		maxTrappedWater += trappedWater
	}

	return maxTrappedWater
}

// time complexity => O(n)
// space complexity => O(1)
func trap3(height []int) int {
	maxTrappedWater := 0
	n := len(height)

	if n < 3 {
		return maxTrappedWater
	}

	maxLeftHeight := height[0]
	maxRightHeight := height[n-1]

	left := 0
	right := n - 1

	// calculate trpped water for each wall height
	for left < right {
		if maxLeftHeight < maxRightHeight {
			// start from left
			left++
			maxLeftHeight = max(maxLeftHeight, height[left])
			trappedWater := maxLeftHeight - height[left]
			maxTrappedWater += trappedWater
			continue
		}

		// start from right
		right--
		maxRightHeight = max(maxRightHeight, height[right])
		trappedWater := maxRightHeight - height[right]
		maxTrappedWater += trappedWater
	}

	return maxTrappedWater
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
