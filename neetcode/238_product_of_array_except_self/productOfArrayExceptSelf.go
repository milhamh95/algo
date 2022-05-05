package product_of_array_except_self

// Option 1 -> Brute force -> O(n ^ 2)
// res = []
// for idx, elem in nums {
//  sum := 1
//	for j := idx; j < len(nums); j++ {
//		if idx == j {
//			continue
//		}
//
//		sum = sum * nums[j
//	}
// res = append(res, sum)
//}

// [1,2,3,4] => [24,12,8,6]
// totalFromLeft = [1, (1*2 = 2), (3*2 = 6), (4*6 = 24)]
// totalFromLeft = push(totalFromLeft, nums[0])
// for i = 1; i < len(nums); i++
//		res = nums[i] * totalFromLeft[i-1]
// 		totalFromLeft = pushEnd(totalFromLeft, res)
//
// totalFromRight = [(1*24 = 12) ,(2*12 = 24),(3*4 = 12), 4]
// for i = len(nums) - 2; i >= 0; i--
//		res = nums[i] * totalFromRight[i+1]
//      totalFromright = pushStart(totalFromRight, res)
// for num in nums
//
// nums = [1,2,3,4]
// totalFromLeft = [1, (1*1 = 1), (1*2 = 1), (3*2 = 6)]
// totalFromRight = [(2*12 = 24) ,(3*4 = 12),(1*4 = 4), 1]
// result = []
// for idx in nums
//		res = append(totalFromLeft[idx] * totalFromright[idx])

// time complexity = O(n)
// space complexity = O(n)
func productExceptSelf1(nums []int) []int {
	length := len(nums)
	result := make([]int, len(nums))

	totalFromLeft := make([]int, len(nums))
	totalFromLeft[0] = 1
	for i := 1; i < length; i++ {
		res := nums[i-1] * totalFromLeft[i-1]
		totalFromLeft[i] = res
	}
	// totalFromLeft = [1,1,2,6]

	totalFromRight := make([]int, len(nums))
	totalFromRight[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		res := nums[i+1] * totalFromRight[i+1]
		totalFromRight[i] = res
	}
	// totalFromRight = [24,12,4,1]

	for i := range nums {
		result[i] = totalFromLeft[i] * totalFromRight[i]
	}

	return result
}

func productExceptSelf2(nums []int) []int {
	length := len(nums)
	result := make([]int, len(nums))

	totalFromLeft := make([]int, len(nums))
	totalFromLeft[0] = 1

	totalFromRight := make([]int, len(nums))
	totalFromRight[length-1] = 1

	// totalFromLeft = [1,1,2,6]
	// totalFromRight = [24,12,4,1]
	for i := 1; i < length; i++ {
		totalFromLeft[i] = nums[i-1] * totalFromLeft[i-1]
		totalFromRight[length-i-1] = nums[length-i] * totalFromRight[length-i]
	}

	for i := range nums {
		result[i] = totalFromLeft[i] * totalFromRight[i]
	}

	return result
}
