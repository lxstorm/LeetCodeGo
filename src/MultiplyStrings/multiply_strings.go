package MultiplyStrings

func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	sums := make([]int, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			mul := int(byte(num1[i])-'0') * int(byte(num2[j]-'0'))
			// place in i+j, i+j+1 index
			newSum := 10*(sums[i+j]) + sums[i+j+1] + mul
			sums[i+j], sums[i+j+1] = newSum%100/10, newSum%10
			carry := newSum / 100
			for k := i + j - 1; k >= 0 && carry != 0; k-- {
				tmp := sums[k] + carry
				sums[k] = tmp % 10
				carry = tmp / 10
			}
		}
	}
	byteSum := []byte{}
	startPos := 0
	for ; startPos < len(sums)-1 && sums[startPos] == 0; startPos++ {
	}
	for i := 0; i < len(sums); i++ {
		byteSum = append(byteSum, byte(sums[i]+int('0')))
	}
	return string(byteSum[startPos:len(byteSum)])
}
