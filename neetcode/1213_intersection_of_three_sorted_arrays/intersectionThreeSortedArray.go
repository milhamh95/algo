package intersectionofthreesortedarrays

func intersection(nums1, nums2, nums3 []int) []int {
	res := []int{}

	ptrNums1 := 0
	ptrNums2 := 0
	ptrNums3 := 0

	for ptrNums1 < len(nums1) &&
		ptrNums2 < len(nums2) &&
		ptrNums3 < len(nums3) {
		v1 := nums1[ptrNums1]
		v2 := nums2[ptrNums2]
		v3 := nums3[ptrNums3]

		if v1 == v2 && v1 == v3 {
			res = append(res, v1)
			ptrNums1++
			ptrNums2++
			ptrNums3++
			continue
		}

		if v1 < v2 || v1 < v3 {
			ptrNums1++
			continue
		}

		if v2 < v1 || v2 < v3 {
			ptrNums2++
			continue
		}

		if v3 < v1 || v3 < v2 {
			ptrNums3++
			continue
		}
	}

	return res
}
