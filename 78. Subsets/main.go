package main

func subsets(nums []int) [][]int {
	result := [][]int{[]int{}}
	for _, num := range nums {
		for _, item := range result {
			result = append(result, append(append([]int{}, item...), num))
		}
	}

	return result
}