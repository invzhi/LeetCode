// 350. Intersection of Two Arrays II

// Given two arrays, write a function to compute their intersection.

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]

// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]

// Note:
// - Each element in the result should appear as many times as it shows in both arrays.
// - The result can be in any order.

// Follow up:
// - What if the given array is already sorted? How would you optimize your algorithm?
// - What if nums1's size is small compared to nums2's size? Which algorithm is better?
// - What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

package leetcode

import "sort"

func intersect(nums1, nums2 []int) []int {
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
			nums = append(nums, nums1[i])
			i++
			j++
		}
	}
	return nums
}
