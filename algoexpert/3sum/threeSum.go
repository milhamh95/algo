package threesum

import "sort"

func ThreeSum(nums []int, target int) [][]int {
	res := [][]int{}

	if len(nums) == 0 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		l := i + 1
		r := len(nums) - 1

		for l < r {
			total := nums[i] + nums[l] + nums[r]

			if total == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				continue
			}

			if total < target {
				l++
				continue
			}

			if total > target {
				r--
				continue
			}
		}
	}

	return res
}
