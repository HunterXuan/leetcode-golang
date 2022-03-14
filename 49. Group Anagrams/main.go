package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)

	group := make(map[string][]string)
	for _, str := range strs {
		keyStr := str
		charArr := strings.Split(keyStr, "")
		sort.Strings(charArr)
		keyStr = strings.Join(charArr, "")
		group[keyStr] = append(group[keyStr], str)
	}

	for _, items := range group {
		result = append(result, items)
	}

	return result
}