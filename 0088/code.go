// 88. Merge Sorted Array

// Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

// Note:
// - The number of elements initialized in nums1 and nums2 are m and n respectively.
// - You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.

// Example:
// Input:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3
// Output: [1,2,2,3,5,6]

package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	lo, hi := 0, m-1
	for _, num := range nums2 {
		p := position(nums1, num, lo, hi)
		lo, hi = p+1, hi+1

		for i := m; i > p; i-- {
			nums1[i] = nums1[i-1]
		}
		nums1[p] = num
		m++
	}
}

func position(nums []int, k int, lo, hi int) int {
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > k {
			if lo == hi {
				break
			}
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
