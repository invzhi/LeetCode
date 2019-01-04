// 863. All Nodes Distance K in Binary Tree

// We are given a binary tree (with root node root), a target node, and an integer value K.
// Return a list of the values of all nodes that have a distance K from the target node.  The answer can be returned in any order.

// Example 1:
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
// Output: [7,4,1]
// Explanation:
// The nodes that are a distance 2 from the target node (with value 5) have values 7, 4, and 1.
// Note that the inputs "root" and "target" are actually TreeNodes.
// The descriptions of the inputs above are just serializations of these objects.

// Note:
// 1. The given tree is non-empty.
// 2. Each node in the tree has unique values 0 <= node.val <= 500.
// 3. The target node is a node in the tree.
// 4. 0 <= K <= 1000.

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root, target *TreeNode, K int) []int {
	var vals []int

	var f func(*TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		if root == target {
			vals = down(vals, K, root)
			return 0
		}

		var d int
		var to *TreeNode

		if d = f(root.Left); d != -1 {
			to = root.Right
		} else if d = f(root.Right); d != -1 {
			to = root.Left
		} else {
			return -1
		}

		if d+1 == K {
			vals = append(vals, root.Val)
		} else if d+1 < K {
			vals = down(vals, K-d-2, to)
		}

		return d + 1
	}
	f(root)

	return vals
}

func down(vals []int, distance int, root *TreeNode) []int {
	if root == nil {
		return vals
	}
	if distance == 0 {
		return append(vals, root.Val)
	}

	vals = down(vals, distance-1, root.Left)
	vals = down(vals, distance-1, root.Right)

	return vals
}
