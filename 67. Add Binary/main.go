package main

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	var result string
	aLen, bLen, idx, left := len(a), len(b), 0, uint8(0)
	for idx < len(b) {
		tmp := (a[aLen - idx - 1] - '0') + (b[bLen - idx - 1] - '0') + left
		if tmp > 1 {
			tmp = tmp - 2
			left = 1
		} else {
			left = 0
		}
		result = string(tmp + '0') + result
		idx = idx + 1
	}

	for idx < len(a) {
		tmp := a[aLen - idx - 1] - '0' + left
		if tmp > 1 {
			tmp = tmp - 2
			left = 1
		} else {
			left = 0
		}
		result = string(tmp + '0') + result
		idx = idx + 1
	}

	if left == 1 {
		result = "1" + result
	}

	return result
}