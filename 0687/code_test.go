package leetcode

import (
	"fmt"
	"testing"
)

func (t *TreeNode) String() string {
	return fmt.Sprint(t.values())
}

func (t *TreeNode) values() []interface{} {
	if t == nil {
		return nil
	}

	vs := []interface{}{t.Val}

	for q := []*TreeNode{t.Left, t.Right}; len(q) > 0; q = q[1:] {
		if q[0] == nil {
			vs = append(vs, nil)
			continue
		}
		vs = append(vs, q[0].Val)
		q = append(q, q[0].Left, q[0].Right)
	}

	i := len(vs) - 1
	for vs[i] == nil {
		i--
	}
	return vs[:i+1]
}

func newTree(vs ...interface{}) *TreeNode {
	if len(vs) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(vs))
	for i := len(vs) - 1; i >= 0; i-- {
		if vs[i] == nil {
			continue
		}
		if v, ok := vs[i].(int); ok {
			nodes[i] = &TreeNode{Val: v}
			if 2*i+1 < len(nodes) {
				nodes[i].Left = nodes[2*i+1]
			}
			if 2*i+2 < len(nodes) {
				nodes[i].Right = nodes[2*i+2]
			}
		}
	}
	return nodes[0]
}

func TestLongestUnivaluePath(t *testing.T) {
	var tests = []struct {
		tree   *TreeNode
		length int
	}{
		{newTree(5, 4, 5, 1, 1, nil, 5), 2},
		{newTree(5, 5, 4, nil, 5, 1, 1), 2},
		{newTree(1, 4, 5, 4, 4, nil, 5), 2},
		{newTree(1, 5, 4, 5, nil, 4, 4), 2},
		{newTree(26, 26, 26), 2},
		{newTree(26, 26, 26, 26, 26), 3},
		{newTree(4, -7, -3, nil, nil, -9, -3, 9, -7, -4, nil, 6, nil, -6, -6, nil, nil, 0, 6, 5, nil, 9, nil, nil, -1, -4, nil, nil, nil, -2), 1},
	}

	for _, tt := range tests {
		length := longestUnivaluePath(tt.tree)
		if length != tt.length {
			t.Errorf("longestUnivaluePath(%v) return %v, want %v", tt.tree, length, tt.length)
		}
	}
}
