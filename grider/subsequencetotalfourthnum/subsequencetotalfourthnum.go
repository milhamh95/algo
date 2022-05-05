package subsequencetotalfourthnum

func subsequencenum(n int) []int {
	nums := []int{1, 2, 3, 4}

	for i := len(nums); i <= n; i++ {
		idx4 := len(nums) - 1
		idx3 := len(nums) - 2
		idx2 := len(nums) - 3
		idx1 := len(nums) - 4

		sum := nums[idx1] + nums[idx2] + nums[idx3] + nums[idx4]
		if i == n {
			return []int{nums[idx1-1], nums[idx2-1], nums[idx3-1], nums[idx4-1]}
		}

		nums = append(nums, sum)
	}

	return []int{}
}
