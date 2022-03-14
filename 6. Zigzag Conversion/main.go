package main

func convert(s string, numRows int) string {
	result := ""
	strArr := make([]string, numRows)
	roundCount := (numRows - 1) * 2
	if roundCount <= 0 {
		return s
	}

	for i := 0; i < len(s); i++ {
		row := i % roundCount
		if row > numRows - 1 {
			row = roundCount - row
		}

		strArr[row] = strArr[row] + string(s[i])
	}

	for _, str := range strArr {
		result = result + str
	}

	return result
}