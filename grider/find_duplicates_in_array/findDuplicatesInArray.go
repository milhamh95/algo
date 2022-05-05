package find_duplicates_in_array

// we can sort the array
// but time complexity will become n log n
// i = 0, i < len(nums) - 1
// if nums[i] == nums[i+1] = push to array

// time complexity => O(n)
// space complexity => O(n)
func findDuplicates(nums []int) []int {
	res := []int{}
	mapNums := map[int]int{}

	for _, v := range nums {
		mapNums[v]++

		ctr := mapNums[v]
		if ctr == 2 {
			res = append(res, v)
		}
	}

	return res
}
