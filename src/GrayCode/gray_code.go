package GrayCode

import "math"

func powWrapper(n int) int {
	return int(math.Pow(float64(2), float64(n)))
}

func grayCode(n int) []int {
	ret := []int{}
	for i := 0; i < powWrapper(n); i++ {
		ret = append(ret, (i>>1)^i)
	}

	return ret

}
