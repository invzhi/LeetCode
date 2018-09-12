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

func TestHasPathSum(t *testing.T) {
	var tests = []struct {
		tree   *TreeNode
		sum    int
		hasSum bool
	}{
		{newTree(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, nil, nil, 1), 22, true},
		{newTree(5, 4, 8, 11, nil, 13, 4, nil, 2, nil, nil, nil, nil, nil, 1), 22, true},
	}

	for _, tt := range tests {
		hasSum := hasPathSum(tt.tree, tt.sum)
		if hasSum != tt.hasSum {
			t.Errorf("hasPathSum(%v, %v) return %v, want %v", tt.tree, tt.sum, hasSum, tt.hasSum)
		}
	}
}
