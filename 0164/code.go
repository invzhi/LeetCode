// 164. Maximum Gap

// Given an unsorted array, find the maximum difference between the successive elements in its sorted form.
// Return 0 if the array contains less than 2 elements.

// Example 1:
// Input: [3,6,9,1]
// Output: 3
// Explanation: The sorted form of the array is [1,3,6,9], either (3,6) or (6,9) has the maximum difference 3.

// Example 2:
// Input: [10]
// Output: 0
// Explanation: The array contains less than 2 elements, therefore return 0.

// Note:
// - You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.
// - Try to solve it in linear time/space.

package leetcode

type bucket struct {
	min, max int
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	min, max := findMinAndMax(nums)
	size := (max - min) / (len(nums) - 1)
	if size < 1 {
		size = 1
	}

	buckets := make([]*bucket, (max-min)/size+1)
	for _, num := range nums {
		i := (num - min) / size
		if buckets[i] == nil {
			buckets[i] = &bucket{min: num, max: num}
			continue
		}
		if num < buckets[i].min {
			buckets[i].min = num
		} else if num > buckets[i].max {
			buckets[i].max = num
		}
	}

	var gap int
	prevMax := buckets[0].max
	for i := 1; i < len(buckets); i++ {
		if buckets[i] != nil {
			if buckets[i].min-prevMax > gap {
				gap = buckets[i].min - prevMax
			}
			prevMax = buckets[i].max
		}
	}
	return gap
}

func findMinAndMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}
	return min, max
}
