package containsduplicate

// time complexity -> O(n)
// space complexity -> O(n)
func containsDuplicate(nums []int) bool {
	mapNums := map[int]bool{}

	for _, v := range nums {
		_, ok := mapNums[v]
		if !ok {
			mapNums[v] = true
			continue
		}

		return true
	}

	return false
}
