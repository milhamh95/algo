package twosumII

// time complex = O(n)
// space complexity = O(n)
func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	mapNums := map[int]int{}

	for i, v1 := range numbers {
		v2 := target - v1

		idx, ok := mapNums[v2]
		if ok {
			return []int{idx + 1, i + 1}
		}

		mapNums[v1] = i
	}

	return []int{}
}

// time complexity = O(n)
// space complexity = O(1)
func twoSum2(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		total := nums[left] + nums[right]
		if total == target {
			return []int{left + 1, right + 1}
		}

		if total < target {
			left++
			continue
		}

		right--

	}

	return []int{}
}

// time complexity = O(n)
// space complexity = O(1)
func twoSum3(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	total := nums[left] + nums[right]

	for total != target {
		if total < target {
			left++
		} else {
			right--
		}

		total = nums[left] + nums[right]
	}

	return []int{left + 1, right + 1}
}
