package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	lengthNums := len(nums)
	if lengthNums == 0 {
		return 0
	}

	ctr, i := 1, 1

	for i < lengthNums {
		if nums[i] != nums[i-1] {
			nums[ctr] = nums[i]
			ctr++
		}
		i++
	}

	return ctr
}
