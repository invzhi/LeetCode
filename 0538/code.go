// 538. Convert BST to Greater Tree

// Given a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.

// Example:
// Input: The root of a Binary Search Tree like this:
//               5
//             /   \
//            2     13
// Output: The root of a Greater Tree like this:
//              18
//             /   \
//           20     13

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	for p := root; p != nil; {
		if p.Right == nil {
			sum += p.Val
			p.Val = sum
			p = p.Left
			continue
		}

		q := p.Right
		for q.Left != nil && q.Left != p {
			q = q.Left
		}

		if q.Left != nil {
			sum += p.Val
			p.Val = sum
			q.Left = nil
			p = p.Left
		} else {
			q.Left = p
			p = p.Right
		}
	}
	return root
}
