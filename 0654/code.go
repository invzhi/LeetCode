// 654. Maximum Binary Tree

// Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:
// 1. The root is the maximum number in the array.
// 2. The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
// 3. The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
// Construct the maximum tree by the given array and output the root node of this tree.

// Example 1:
// Input: [3,2,1,6,0,5]
// Output: return the tree root node representing the following tree:
//       6
//     /   \
//    3     5
//     \    /
//      2  0
//        \
//         1

// Note: The size of the given array will be in the range [1,1000].

package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var max, maxi int
	for i, num := range nums {
		if num > max {
			max = num
			maxi = i
		}
	}

	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:maxi]),
		Right: constructMaximumBinaryTree(nums[maxi+1:]),
	}
}
