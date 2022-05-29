package longest_increasing_subsequence

import (
	"fmt"
	"sort"
)

func lengthOfLIS1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lenCounterSub := []int{1}

	// maximum length of increasing subsequence
	maxLongestSub := 1

	for i := 1; i < len(nums); i++ {
		currentLongestSub := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				currentLongestSub = max(currentLongestSub, lenCounterSub[j]+1)
			}
		}

		lenCounterSub = append(lenCounterSub, currentLongestSub)
		maxLongestSub = max(maxLongestSub, currentLongestSub)
	}

	fmt.Println("========  ========")
	fmt.Printf("%+v\n", lenCounterSub)
	fmt.Println("=================")

	return maxLongestSub
}

func lengthOfLIS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// length of inceasing subsequence each index
	lenCounterSub := make([]int, len(nums))
	for i := range nums {
		lenCounterSub[i] = 1
	}

	// maximum length of increasing subsequence
	maxLongestSub := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				lenCounterSub[i] = max(lenCounterSub[i], lenCounterSub[j]+1)
			}
		}

		maxLongestSub = max(maxLongestSub, lenCounterSub[i])
	}

	return maxLongestSub
}

func lengthOfLIS3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	subseq := []int{}
	for _, v := range nums {
		idx := sort.SearchInts(subseq, v)
		if idx == len(subseq) {
			subseq = append(subseq, v)
			continue
		} else {
			subseq[idx] = v
		}
	}

	return len(subseq)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
