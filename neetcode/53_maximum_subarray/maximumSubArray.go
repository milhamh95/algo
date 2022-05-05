package maximumsubarray

// maxSum = 0
// for i := 0
// 		tmpSum = nums[i]
// 			for j := i + 1; j < n; j++
//				tmpSum = tmpSum + nums[j]
//				maxSum = max(tmpSum)
// O n^2

// maxSum = -99999
// tmpSum = nums[0]
// for i := 1; i < len(n); i++
// 		tmpSum = tmpSum + nums[i]
//      tmpSum = max(tmpSum, nums[i])
//      maxSum = max(tmpSum, max)
// [-2,1,-3,4,-1,2,1,-5,4]
func maxSubArray(nums []int) int {
	tmpSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		tmpSum = tmpSum + nums[i]
		tmpSum = max(tmpSum, nums[i])
		maxSum = max(maxSum, tmpSum)
	}

	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
