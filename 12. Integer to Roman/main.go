package main

func intToRoman(num int) string {
	var result string

	intRomanArr := []struct{
		Number int
		Roman string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, val := range intRomanArr {
		for num > 0 {
			if num >= val.Number {
				num = num - val.Number
				result = result + val.Roman
			} else {
				break
			}
		}
	}

	return result
}