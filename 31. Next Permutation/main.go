package main

func divide(dividend int, divisor int) int {
	maxCount := 1 << 31

	flag := 1
	if dividend < 0 {
		flag = -flag
		dividend = -dividend
	}

	if divisor < 0 {
		flag = -flag
		divisor = -divisor
	}

	count := 0
	unitCount, unit := 1, divisor
	for dividend >= unit {
		dividend = dividend - unit
		count = count + unitCount
		if count >= maxCount {
			if flag > 0 {
				return maxCount - 1
			} else {
				return -maxCount
			}
		}

		unit = unit << 1
		unitCount = unitCount << 1
	}

	unit = unit >> 1
	unitCount = unitCount >> 1
	for unit > 0 {
		for dividend >= unit {
			dividend = dividend - unit
			count = count + unitCount
			if count >= maxCount {
				if flag > 0 {
					return maxCount - 1
				} else {
					return -maxCount
				}
			}
		}

		unit = unit >> 1
		unitCount = unitCount >> 1
	}

	return flag * count
}