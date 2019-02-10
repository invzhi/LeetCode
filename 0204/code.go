// 204. Count Primes

// Count the number of prime numbers less than a non-negative number, n.

// Example:
// Input: 10
// Output: 4
// Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.

package leetcode

func countPrimes(n int) int {
	primes := make([]bool, n)
	for i := 2; i < n; i++ {
		primes[i] = true
	}

	for i := 2; i*i < n; i++ {
		if primes[i] {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}

	var cnt int
	for i := 2; i < n; i++ {
		if primes[i] {
			cnt++
		}
	}
	return cnt
}
