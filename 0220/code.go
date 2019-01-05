// 20. Contains Duplicate III

// Given an array of integers, find out whether there are two distinct indices i and j in the array such that the absolute difference between nums[i] and nums[j] is at most t and the absolute difference between i and j is at most k.

// Example 1:
// Input: nums = [1,2,3,1], k = 3, t = 0
// Output: true

// Example 2:
// Input: nums = [1,0,1,1], k = 1, t = 2
// Output: true

// Example 3:
// Input: nums = [1,5,9,1,5,9], k = 2, t = 3
// Output: false

package leetcode

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	bucket := make(map[int]int, k+1)
	for i, num := range nums {
		if t == 0 {
			if j, ok := bucket[num]; ok && i-j <= k {
				return true
			}
			bucket[num] = i
			continue
		}

		index := num / t
		for j := index - 1; j <= index+1; j++ {
			if v, ok := bucket[j]; ok && d(num, v) <= t {
				return true
			}
		}

		bucket[index] = num
		if len(bucket) > k {
			delete(bucket, nums[i-k]/t)
		}
	}
	return false
}

func d(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
