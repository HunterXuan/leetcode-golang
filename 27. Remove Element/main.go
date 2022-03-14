package main

func removeElement(nums []int, val int) int {
	pointer := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[pointer] = nums[i]
			pointer = pointer + 1
			count = count + 1
		}
	}

	return count
}