package main

func jump(nums []int) int {
	numsLen := len(nums)
	minStepsArr := make([]int, len(nums))

	for i := numsLen - 2; i >= 0; i-- {
		minStep := numsLen
		for j := 1; j <= nums[i]; j++ {
			curIdx := i + j
			if curIdx < numsLen && 1 + minStepsArr[curIdx] < minStep {
				minStep = 1 + minStepsArr[curIdx]
			}
		}

		minStepsArr[i] = minStep
	}

	return minStepsArr[0]
}
