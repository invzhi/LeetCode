// 4. Median of Two Sorted Arrays

// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
// You may assume nums1 and nums2 cannot be both empty.

// Example 1:
// nums1 = [1, 3]
// nums2 = [2]
// The median is 2.0

// Example 2:
// nums1 = [1, 2]
// nums2 = [3, 4]
// The median is (2 + 3)/2 = 2.5

package leetcode

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	half := (len(nums1) + len(nums2) + 1) / 2

	var median float64
	lo, hi := 0, len(nums1)
	for lo <= hi {
		i := lo + (hi-lo)/2
		j := half - i

		if i > 0 && nums1[i-1] > nums2[j] {
			hi = i - 1
		} else if i < len(nums1) && nums2[j-1] > nums1[i] {
			lo = i + 1
		} else {
			var l int
			if i == 0 {
				l = nums2[j-1]
			} else if j == 0 {
				l = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					l = nums1[i-1]
				} else {
					l = nums2[j-1]
				}
			}

			if (len(nums1)+len(nums2))&1 == 1 {
				median = float64(l)
				break
			}

			var r int
			if i == len(nums1) {
				r = nums2[j]
			} else if j == len(nums2) {
				r = nums1[i]
			} else {
				if nums1[i] < nums2[j] {
					r = nums1[i]
				} else {
					r = nums2[j]
				}
			}

			median = float64(l+r) / 2
			break
		}
	}
	return median
}
