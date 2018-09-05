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

func intersection(nums1, nums2 []int) []int {
	m := make(map[int]struct{})
	for _, num := range nums1 {
		m[num] = struct{}{}
	}

	nums := make([]int, 0)
	for _, num := range nums2 {
		if _, ok := m[num]; ok {
			nums = append(nums, num)
			delete(m, num)
		}
	}
	return nums
}
