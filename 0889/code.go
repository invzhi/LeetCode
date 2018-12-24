// 889. Construct Binary Tree from Preorder and Postorder Traversal

// Return any binary tree that matches the given preorder and postorder traversals.
// Values in the traversals pre and post are distinct positive integers.

// Example 1:
// Input: pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
// Output: [1,2,3,4,5,6,7]

// Note:
// - 1 <= pre.length == post.length <= 30
// - pre[] and post[] are both permutations of 1, 2, ..., pre.length.
// - It is guaranteed an answer exists. If there exists multiple answers, you can return any of them.

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(pre, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	if len(pre) == 1 {
		return &TreeNode{Val: pre[0]}
	}

	n := len(post)
	i := findIndex(post, pre[1])
	return &TreeNode{
		Val:   pre[0],
		Left:  constructFromPrePost(pre[1:i+2], post[:i+1]),
		Right: constructFromPrePost(pre[i+2:], post[i+1:n-1]),
	}
}

func findIndex(order []int, v int) int {
	var i int
	for order[i] != v {
		i++
	}
	return i
}
