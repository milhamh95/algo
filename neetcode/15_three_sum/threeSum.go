package threesum

import "sort"

// O (n log n) + O (n ^ 2) -> O(n ^ 2)
func threeSum(nums []int) [][]int {
	res := [][]int{}

	if len(nums) == 0 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := n1 + nums[l] + nums[r]

			if sum > 0 {
				r--
				continue
			}

			if sum < 0 {
				l++
				continue
			}

			if sum == 0 {
				res = append(res, []int{n1, nums[l], nums[r]})

				// check if num l equal with num l + 1
				// if equal, found duplicate. increase l
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				// check if num r equal with num r - 1
				// if equal, found duplicate. decrease r
				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
				continue
			}
		}
	}

	return res
}
