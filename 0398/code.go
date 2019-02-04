// 398. Random Pick Index

// Given an array of integers with possible duplicates, randomly output the index of a given target number. You can assume that the given target number must exist in the array.

// Note:
// The array size can be very large. Solution that uses too much extra space will not pass the judge.

// Example:
// int[] nums = new int[] {1,2,3,3,3};
// Solution solution = new Solution(nums);
// solution.pick(3); // pick(3) should return either index 2, 3, or 4 randomly. Each index should have equal probability of returning.
// solution.pick(1); // pick(1) should return 0. Since in the array only nums[0] is equal to 1.

package leetcode

import "math/rand"

// Solution is the solution for the reservoir sampling algorithm.
type Solution struct {
	nums []int
}

// Constructor accepts a slice of integers.
func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

// Pick randomly output the index of a given target number.
func (s *Solution) Pick(target int) int {
	index := -1

	var j int
	for i, num := range s.nums {
		if num == target {
			if rand.Intn(j+1) == 0 {
				index = i
			}
			j++
		}
	}

	return index
}
