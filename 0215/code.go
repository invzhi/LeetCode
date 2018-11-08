// 215. Kth Largest Element in an Array

// Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

// Example 1:
// Input: [3,2,1,5,6,4] and k = 2
// Output: 5

// Example 2:
// Input: [3,2,3,1,2,4,5,5,6] and k = 4
// Output: 4

// Note:
// You may assume k is always valid, 1 ≤ k ≤ array's length.

package leetcode

func findKthLargest(nums []int, k int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := partition(nums, lo, hi)
		if k-1 < mid {
			hi = mid - 1
		} else if k-1 > mid {
			lo = mid + 1
		} else {
			break
		}
	}
	return nums[k-1]
}

func partition(nums []int, lo, hi int) int {
	pivot := nums[lo]
	i, j := lo, hi+1
	for {
		for i+1 <= hi && nums[i+1] > pivot {
			i++
		}
		i++
		for nums[j-1] < pivot {
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
