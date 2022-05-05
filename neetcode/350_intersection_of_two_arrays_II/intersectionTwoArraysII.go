package intersectionoftwoarraysii

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}

	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}

	mapNums1 := map[int]int{}

	for _, v := range nums1 {
		mapNums1[v]++
	}

	for _, v := range nums2 {
		counter, ok := mapNums1[v]
		if ok && counter > 0 {
			res = append(res, v)
			mapNums1[v]--
		}
	}

	return res
}
