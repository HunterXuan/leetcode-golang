package main

func romanToInt(s string) int {
	var num int

	romanIntMap := map[string]int{
		"M": 1000,
		"CM": 900,
		"D": 500,
		"CD": 400,
		"C": 100,
		"XC": 90,
		"L": 50,
		"XL": 40,
		"X": 10,
		"IX": 9,
		"V": 5,
		"IV": 4,
		"I": 1,
	}

	strLen := len(s)
	pointer := 0
	for pointer < strLen {
		if pointer + 1 < strLen {
			tmp := s[pointer:pointer+2]
			if val, ok := romanIntMap[tmp]; ok {
				num = num + val
				pointer = pointer + 2
				continue
			}
		}

		tmp := s[pointer]
		if val, ok := romanIntMap[string(tmp)]; ok {
			num = num + val
			pointer = pointer + 1
			continue
		} else {
			break
		}
	}

	return num
}