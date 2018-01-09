package leetcode

import "testing"

func insertLeaf(root, leaf *TreeNode) {
	for root != nil {
		switch {
		case root.Val < leaf.Val:
			if root.Right == nil {
				root.Right = leaf
				return
			}
			root = root.Right
		case root.Val > leaf.Val:
			if root.Left == nil {
				root.Left = leaf
				return
			}
			root = root.Left
		}
	}
}

func buildTree(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	l := len(nums)

	for i := 1; i < l; i++ {
		leaf := &TreeNode{Val: nums[i]}
		insertLeaf(root, leaf)
	}
	return root
}

func TestFindTarget(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		res    bool
	}{
		{[]int{5, 3, 6, 2, 4, 7}, 9, true},
		{[]int{5, 3, 6, 2, 4, 7}, 28, false},
		{[]int{5, 3}, 8, true},
		{nil, 28, false},
	}

	for _, tt := range tests {
		tree := buildTree(tt.nums)
		res := findTarget(tree, tt.target)
		if res != tt.res {
			t.Errorf("findTarget(%v, %v) return %v, want %v", tt.nums, tt.target, res, tt.res)
		}
	}
}
