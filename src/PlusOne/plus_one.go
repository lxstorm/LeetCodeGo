package PlusOne

func plusOne(digits []int) []int {
	carry := 1
	output := make([]int, len(digits)+1)
	for n := len(digits) - 1; n >= 0; n-- {
		curDigit := (digits[n] + carry) % 10
		carry = (digits[n] + carry) / 10
		output[n+1] = curDigit
	}
	if carry == 0 {
		output = output[1:]
	} else {
		output[0] = carry
	}

	return output
}
