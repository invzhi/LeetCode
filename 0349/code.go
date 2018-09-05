// 349. Intersection of Two Arrays

// Given two arrays, write a function to compute their intersection.

// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]

// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]

// Note:
// - Each element in the result must be unique.
// - The result can be in any order.

package leetcode

import "sort"

func intersection(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var (
		i, j   int
		l1, l2 = len(nums1), len(nums2)
		nums   = make([]int, 0)
	)
	for i < l1 && j < l2 {
		switch {
		case nums1[i] < nums2[j]:
			i++
		case nums1[i] > nums2[j]:
			j++
		default:
			if l, v := len(nums), nums1[i]; l == 0 || nums[l-1] != v {
				nums = append(nums, v)
			}
			i++
			j++
		}
	}
	return nums
}
