// 324. Wiggle Sort II

// Given an unsorted array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....

// Example 1:
// Input: nums = [1, 5, 1, 1, 6, 4]
// Output: One possible answer is [1, 4, 1, 5, 1, 6].

// Example 2:
// Input: nums = [1, 3, 2, 2, 3, 1]
// Output: One possible answer is [2, 3, 1, 3, 1, 2].

// Note:
// You may assume all input has valid answer.

// Follow Up:
// Can you do it in O(n) time and/or in-place with O(1) extra space?

package leetcode

func wiggleSort(nums []int) {
	n := len(nums) | 1

	mid := findMiddle(nums)

	gt, lt := 1, (len(nums)-1)&(^1)
	for i := 1; i&1 == 1 || i <= lt; {
		if i == n {
			i = 0
		}

		if nums[i] < mid {
			nums[i], nums[lt] = nums[lt], nums[i]
			lt -= 2
		} else if nums[i] > mid {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt += 2
			i += 2
		} else {
			i += 2
		}
	}
}

func findMiddle(nums []int) int {
	lo, hi := 0, len(nums)-1
	mi := hi / 2
	for lo < hi {
		mid := partition(nums, lo, hi)
		if mi < mid {
			hi = mid - 1
		} else if mi > mid {
			lo = mid + 1
		} else {
			break
		}
	}
	return nums[mi]
}

func partition(nums []int, lo, hi int) int {
	pivot := nums[lo]
	i, j := lo, hi+1
	for {
		for i+1 <= hi && nums[i+1] < pivot {
			i++
		}
		i++
		for nums[j-1] > pivot {
			j--
		}
		j--

		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}
