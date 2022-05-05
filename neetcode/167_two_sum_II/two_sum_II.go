package twosumII

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
