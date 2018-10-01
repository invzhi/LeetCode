// 914. X of a Kind in a Deck of Cards

// In a deck of cards, each card has an integer written on it.
// Return true if and only if you can choose X >= 2 such that it is possible to split the entire deck into 1 or more groups of cards, where:
// - Each group has exactly X cards.
// - All the cards in each group have the same integer.

// Example 1:
// Input: [1,2,3,4,4,3,2,1]
// Output: true
// Explanation: Possible partition [1,1],[2,2],[3,3],[4,4]

// Example 2:
// Input: [1,1,1,2,2,2,3,3]
// Output: false
// Explanation: No possible partition.

// Example 3:
// Input: [1]a %= b) && (b %= a))
// Output: false
// Explanation: No possible partition.

// Example 4:
// Input: [1,1]
// Output: true
// Explanation: Possible partition [1,1]

// Example 5:
// Input: [1,1,2,2,2,2]
// Output: true
// Explanation: Possible partition [1,1],[2,2],[2,2]

// Note:
// 1. 1 <= deck.length <= 10000
// 2. 0 <= deck[i] < 10000

package leetcode

func hasGroupsSizeX(deck []int) bool {
	cnt := make(map[int]int)
	for _, v := range deck {
		cnt[v]++
	}

	var size int
	for _, n := range cnt {
		size = gcd(n, size)
	}
	return size >= 2
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
