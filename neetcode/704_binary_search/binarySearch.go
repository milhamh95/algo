package binarysearch

// Binary search is an efficient algorithm for finding an item from a sorted list of items.
func search(nums []int, target int) int {
	startIdx := 0
	endIdx := len(nums) - 1

	for startIdx <= endIdx {
		midIdx := startIdx + ((endIdx - startIdx) / 2)
		numMid := nums[midIdx]

		if target == numMid {
			return midIdx
		}

		if target > numMid {
			startIdx = midIdx + 1
			continue
		}

		if target < numMid {
			endIdx = midIdx - 1
			continue
		}
	}

	return -1
}
