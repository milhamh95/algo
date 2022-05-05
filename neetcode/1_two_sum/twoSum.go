package two_sum

func twoSum(nums []int, target int) []int {
	// key = v2
	// value = index
	mapNums := map[int]int{}
	for i, v1 := range nums {
		v2 := target - v1
		idx, ok := mapNums[v2]
		if ok {
			return []int{idx, i}
		}

		mapNums[v1] = i
	}

	return []int{}
}
