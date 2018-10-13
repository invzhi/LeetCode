// 46. Permutations

// Given a collection of distinct integers, return all possible permutations.

// Example:
// Input: [1,2,3]
// Output:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

package leetcode

func permute(nums []int) [][]int {
	permutations := [][]int{nums}
	for i := 0; i+1 < len(nums); i++ {
		instances := permutations
		for j := i + 1; j < len(nums); j++ {
			for _, instance := range instances {
				permutation := make([]int, len(instance))
				copy(permutation, instance)
				permutation[i], permutation[j] = permutation[j], permutation[i]
				permutations = append(permutations, permutation)
			}
		}
	}
	return permutations
}
