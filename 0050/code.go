// 50. Pow(x, n)

// Implement pow(x, n), which calculates x raised to the power n (xn).

// Example 1:
// Input: 2.00000, 10
// Output: 1024.00000

// Example 2:
// Input: 2.10000, 3
// Output: 9.26100

// Example 3:
// Input: 2.00000, -2
// Output: 0.25000
// Explanation: 2-2 = 1/22 = 1/4 = 0.25

// Note:
// - -100.0 < x < 100.0
// - n is a 32-bit signed integer, within the range [−231, 231 − 1]

package leetcode

func myPow(x float64, n int) float64 {
	if x == 0 || x == 1 {
		return x
	}

	var (
		xn float64 = 1
		n_ int64   = int64(n)
	)
	if n < 0 {
		n_ = -n_
	}

	for n_ > 0 {
		if n_&1 == 1 {
			xn *= x
		}
		x *= x
		n_ >>= 1
	}

	if n < 0 {
		xn = 1 / xn
	}
	return xn
}
