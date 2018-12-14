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

	trees := make([][][]*TreeNode, n-1)
	for i := range trees {
		trees[i] = make([][]*TreeNode, n-1-i)
	}
	for d := 1; d < n; d++ {
		for i := 1; i+d <= n; i++ {
			generate(trees, i, i+d)
		}
	}
	return get(trees, 1, n)
}

func generate(trees [][][]*TreeNode, left, right int) {
	var t []*TreeNode
	for i := left; i <= right; i++ {
		ll := get(trees, left, i-1)
		rr := get(trees, i+1, right)
		for _, l := range ll {
			for _, r := range rr {
				t = append(t, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	trees[right-left-1][left-1] = t
}

func get(trees [][][]*TreeNode, left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}
	if left == right {
		return []*TreeNode{{Val: left}}
	}
	return trees[right-left-1][left-1]
}
