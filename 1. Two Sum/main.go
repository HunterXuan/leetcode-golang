package main

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for idx, num := range nums {
		another := target - num
		if pair, ok := tmp[another]; ok {
			return []int{pair, idx}
		} else {
			tmp[num] = idx
		}
	}

	return []int{}
}
