// 95. Unique Binary Search Trees II

// Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

// Example:
// Input: 3
// Output:
// [
//   [1,null,3,2],
//   [3,2,null,1],
//   [3,1,null,null,2],
//   [2,1,3],
//   [1,null,2,null,3]
// ]
// Explanation:
// The above output corresponds to the 5 unique BST's shown below:
//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}

func generate(i, j int) []*TreeNode {
	if i > j {
		return []*TreeNode{nil}
	}

	var trees []*TreeNode
	for k := i; k <= j; k++ {
		ll := generate(i, k-1)
		rr := generate(k+1, j)
		for _, l := range ll {
			for _, r := range rr {
				trees = append(trees, &TreeNode{
					Val:   k,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return trees
}
