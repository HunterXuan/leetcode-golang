package main

func maxSubArray(nums []int) int {
	numsLen := len(nums)
	maxSumArr := nums
	maxSum := nums[numsLen - 1]
	for idx := numsLen - 2; idx >= 0; idx-- {
		if maxSumArr[idx+1] > 0 {
			maxSumArr[idx] = nums[idx] + maxSumArr[idx+1]
		}

		if maxSumArr[idx] > maxSum {
			maxSum = maxSumArr[idx]
		}
	}

	return maxSum
}