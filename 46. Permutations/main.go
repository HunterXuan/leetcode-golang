package main

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		v := make([]int, len(nums))
		copy(v, nums)
		v[i], v[0] = v[0], v[i]

		next := permute(v[1:])
		for _, arr := range next {
			result = append(result, append([]int{nums[i]}, arr...))
		}
	}

	return result
}
