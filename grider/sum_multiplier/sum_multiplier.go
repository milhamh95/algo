package sum_multiplier

import "sort"

// time complexity = O(n ^ 2)
// space complexity = O(1)
func sumMultliplier(nums []int, minTarget int) bool {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}

	totalSum *= 2

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			currentTotal := nums[i] * nums[j]

			if currentTotal >= minTarget {
				return true
			}
		}
	}

	return false
}

// time complexity = O(n log n)
// space complexity = O(1)
func sumMultiplier2(nums []int, mainTarget int) bool {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}

	totalSum *= 2

	//sort.Slice(nums, func(i, j int) bool {
	//	return nums[i] < nums[j]
	//})

	sort.Ints(nums)

	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left] == nums[right] {
			left++
			right--
			continue
		}
		curTotal := nums[left] * nums[right]

		if curTotal > mainTarget {
			return true
		}

		if curTotal < mainTarget {
			left++
			continue
		}

		right--
	}

	return false
}
