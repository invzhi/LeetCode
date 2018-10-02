// 905. Sort Array By Parity

// Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.
// You may return any answer array that satisfies this condition.

// Example 1:
// Input: [3,1,2,4]
// Output: [2,4,3,1]
// The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

// Note:
// 1. 1 <= A.length <= 5000
// 2. 0 <= A[i] <= 5000

package leetcode

func sortArrayByParity(A []int) []int {
	even, odd := 0, len(A)-1
	for even <= odd {
		switch A[even] & 1 {
		case 0:
			even++
		case 1:
			A[even], A[odd] = A[odd], A[even]
			odd--
		}
	}
	return A
}
