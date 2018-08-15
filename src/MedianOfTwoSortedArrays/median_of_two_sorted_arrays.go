package MedianOfTwoSortedArrays

import (
	"math"
)

func findKthNum(nums1 []int, nums2 []int, k int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findKthNum(nums2, nums1, k)
	}
	if m == 0 {
		return float64(nums2[k-1])
	}
	if k == 1 {
		return math.Min(float64(nums1[0]), float64(nums2[0]))
	}
	// cut nums1 index to 0 ~ i-1 and i ~ m - 1
	// cut nums2 index to 0 ~ j-1 and j ~ n - 1
	i := int(math.Min(float64(m), float64(k/2)))
	j := k - i
	if nums1[i-1] > nums2[j-1] {
		return findKthNum(nums1, nums2[j:], k-j)
	} else if nums1[i-1] < nums2[j-1] {
		return findKthNum(nums1[i:], nums2, k-i)
	}
	return float64(nums1[i-1])

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if (m+n)%2 == 1 {
		return findKthNum(nums1, nums2, (m+n)/2+1)
	}
	return (findKthNum(nums1, nums2, (m+n)/2) +
		findKthNum(nums1, nums2, (m+n)/2+1)) / 2.0
}

// not recommend, too trivial
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	m, n := len(nums1), len(nums2)
// 	if m > n {
// 		nums1, nums2, m, n = nums2, nums1, n, m
// 	}
// 	// cut nums1 index to 0 ~ i-1 and i ~ m - 1
// 	// cut nums2 index to 0 ~ j-1 and j ~ n - 1
// 	// left part is (m + n + 1) / 2
// 	imin, imax, halfLen := 0, m, (m+n+1)/2
// 	for imin <= imax {
// 		i := (imin + imax) / 2
// 		j := halfLen - i
// 		// i is too large
// 		if i > imin && nums1[i-1] > nums2[j] {
// 			imax = i - 1
// 		} else if i < imax && nums2[j-1] > nums1[i] {
// 			imin = i + 1
// 		} else {
// 			var maxOfLeft float64
// 			if i == 0 {
// 				maxOfLeft = float64(nums2[j-1])
// 			} else if j == 0 {
// 				maxOfLeft = float64(nums1[i-1])
// 			} else {
// 				maxOfLeft = math.Max(float64(nums1[i-1]),
// 					float64(nums2[j-1]),
// 				)
// 			}
// 			if (m+n)%2 == 1 {
// 				return maxOfLeft
// 			}

// 			var minOfRight float64
// 			if i == m {
// 				minOfRight = float64(nums2[j])
// 			} else if j == n {
// 				minOfRight = float64(nums1[i])
// 			} else {
// 				minOfRight = math.Min(float64(nums1[i]), float64(nums2[j]))
// 			}

// 			return (maxOfLeft + minOfRight) / 2.0

// 		}

// 	}
// 	return -1.0
// }
