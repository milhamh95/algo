package intersectionoftwoarrays

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}

	mapNums1 := map[int]bool{}

	for _, v := range nums1 {
		mapNums1[v] = true
	}

	for _, v := range nums2 {
		num := mapNums1[v]

		if num {
			delete(mapNums1, v)
			res = append(res, v)
		}
	}

	return res
}
