package leetcode

import (
	"fmt"
	"testing"
)

func (t *TreeNode) String() string {
	return fmt.Sprintf("tree(%v)", t.values())
}

func (t *TreeNode) values() []int {
	if t == nil {
		return nil
	}
	vs := []int{t.Val}
	vs = append(vs, t.Left.values()...)
	vs = append(vs, t.Right.values()...)
	return vs
}

func (t *TreeNode) insert(n int) {
	for t != nil {
		switch {
		case t.Val < n:
			if t.Right == nil {
				t.Right = &TreeNode{Val: n}
				return
			}
			t = t.Right
		case t.Val > n:
			if t.Left == nil {
				t.Left = &TreeNode{Val: n}
				return
			}
			t = t.Left
		default:
			return
		}
	}
}

func newTree(nums ...int) *TreeNode {
	if nums == nil {
		return nil
	}

	root := TreeNode{Val: nums[0]}
	l := len(nums)

	for i := 1; i < l; i++ {
		root.insert(nums[i])
	}
	return &root
}

func TestFindTarget(t *testing.T) {
	var tests = []struct {
		tree   *TreeNode
		target int
		exist  bool
	}{
		{newTree(5, 3, 6, 2, 4, 7), 9, true},
		{newTree(5, 3, 6, 2, 4, 7), 4, false},
		{newTree(5, 3), 8, true},
		{newTree(), 2, false},
		{nil, 2, false},
	}

	for _, tt := range tests {
		exist := findTarget(tt.tree, tt.target)
		if exist != tt.exist {
			t.Errorf("findTarget(%v, %v) return %v, want %v", tt.tree, tt.target, exist, tt.exist)
		}
	}
}
