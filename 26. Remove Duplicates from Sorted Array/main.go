package main

func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	} else if numsLen == 1 {
		return 1
	}

	count := 1
	pointer := 1
	for i := 1; i < numsLen; i++ {
		if nums[i] != nums[i-1] {
			nums[pointer] = nums[i]
			pointer = pointer + 1
			count = count + 1
		}
	}

	return count
}