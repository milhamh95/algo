package container_with_most_water

func maxArea(height []int) int {
	area := 0

	l := 0
	r := len(height) - 1

	for l < r {
		length := r - l
		tmpArea := length * min(height[l], height[r])
		area = max(area, tmpArea)

		if height[l] < height[r] {
			l++
			continue
		}

		if height[l] > height[r] {
			r--
			continue
		}

		if height[l] == height[r] {
			l++
			r--
		}
	}

	return area
}

func maxArea1(height []int) int {
	area := 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			length := j - i
			tmpArea := length * min(height[i], height[j])
			area = max(area, tmpArea)
		}
	}

	return area
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
