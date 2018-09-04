package ContainerWithMostWater

// move the short one means from i-short to i-long do not need concern
func maxArea(height []int) int {
	maxSize := 0
	p := 0
	q := len(height) - 1
	for p < q {
		pHeight := height[p]
		qHeight := height[q]
		var curSize int
		if pHeight > qHeight {
			curSize = (q - p) * qHeight
			q--
		} else {
			curSize = (q - p) * pHeight
			p++
		}
		if curSize > maxSize {
			maxSize = curSize
		}
	}

	return maxSize

}
