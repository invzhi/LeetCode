// 167. Two Sum II - Input array is sorted

// Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.
// The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.

// You may assume that each input would have exactly one solution and you may not use the same element twice.
// Input: numbers={2, 7, 11, 15}, target=9
// Output: index1=1, index2=2

package leetcode

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)

	for i, num := range numbers {
		if j, ok := m[target-num]; ok {
			return []int{j + 1, i + 1}
		}
		m[num] = i
	}
	return nil
}

func twoSum1(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]
		switch {
		case sum < target:
			l++
		case sum > target:
			r--
		default:
			return []int{l + 1, r + 1}
		}
	}
	return nil
}
