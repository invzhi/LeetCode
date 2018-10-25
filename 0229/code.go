// 229. Majority Element II

// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
// Note: The algorithm should run in linear time and in O(1) space.

// Example 1:
// Input: [3,2,3]
// Output: [3]

// Example 2:
// Input: [1,1,1,3,3,2,2,2]
// Output: [1,2]

package leetcode

func majorityElement(nums []int) []int {
	m1, m2 := vote2(nums)
	var c1, c2 int
	for _, num := range nums {
		switch num {
		case m1:
			c1++
		case m2:
			c2++
		}
	}

	var majority []int
	if c1 > len(nums)/3 {
		majority = append(majority, m1)
	}
	if c2 > len(nums)/3 {
		majority = append(majority, m2)
	}
	return majority
}

func vote2(nums []int) (int, int) {
	var m1, m2 int
	var c1, c2 int
	for _, num := range nums {
		switch num {
		case m1:
			c1++
		case m2:
			c2++
		default:
			if c1 == 0 {
				m1, c1 = num, 1
			} else if c2 == 0 {
				m2, c2 = num, 1
			} else {
				c1--
				c2--
			}
		}
	}
	return m1, m2
}
