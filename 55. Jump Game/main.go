package main

func canJump(nums []int) bool {
	numsLen := len(nums)
	canJumpArr := make([]bool, len(nums))

	canJumpArr[numsLen - 1] = true
	for i := numsLen - 2; i >= 0; i-- {
		for j := 1; j <= nums[i]; j++ {
			if i + j < numsLen && canJumpArr[i + j] == true {
				canJumpArr[i] = true
				break
			}
		}
	}

	return canJumpArr[0]
}