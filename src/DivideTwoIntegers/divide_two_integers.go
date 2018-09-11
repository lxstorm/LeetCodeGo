package DivideTwoIntegers

import "math"

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func divide(dividend int, divisor int) int {
	if divisor == 0 || (dividend == math.MinInt32 && divisor == -1) {
		return int(math.MaxInt32)
	}
	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}
	divisor = abs(divisor)
	dividend = abs(dividend)
	res := 0

	for dividend >= divisor {
		tmp, multiple := divisor, 1
		for dividend >= (tmp << 1) {
			multiple = multiple << 1
			tmp = tmp << 1
		}
		res += multiple
		dividend = dividend - tmp
	}

	return res * sign

}
