// 728. Self Dividing Numbers

// A self-dividing number is a number that is divisible by every digit it contains.
// For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.
// Also, a self-dividing number is not allowed to contain the digit zero.
// Given a lower and upper number bound, output a list of every possible self dividing number, including the bounds if possible.

// Example 1:
// Input:
// left = 1, right = 22
// Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]

// Note: The boundaries of each input argument are 1 <= left <= right <= 10000.

package leetcode

func selfDividingNumbers(left int, right int) []int {
	var nums []int
	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) {
			nums = append(nums, i)
		}
	}
	return nums
}

func isSelfDividingNumber(num int) bool {
	for n := num; n > 0; n /= 10 {
		d := n % 10
		if d == 0 || num%d > 0 {
			return false
		}
	}
	return true
}
