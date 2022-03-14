package main

func letterCombinations(digits string) []string {
	var result []string

	if len(digits) == 0 {
		return result
	}

	digitCharMap := map[uint8]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	for i := 0; i < len(digits); i++ {
		mapStr, _ := digitCharMap[digits[i]]
		result = recursiveAppendString(result, mapStr)
	}

	return result
}

func recursiveAppendString(strList []string, appendStr string) []string {
	var result []string

	if len(strList) == 0 {
		for i := 0; i < len(appendStr); i++ {
			result = append(result, string(appendStr[i]))
		}
		return result
	}

	for i := 0; i < len(appendStr); i++ {
		for _, str := range strList {
			result = append(result, str + string(appendStr[i]))
		}
	}

	return result
}