package leetcode

import (
	"fmt"
	"reflect"
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

func TestLevelOrder(t *testing.T) {
	var tests = []struct {
		tree  *TreeNode
		order [][]int
	}{
		{
			newTree(3, 9, 20, nil, nil, 15, 7),
			[][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			newTree(1, 2, 2, 3, nil, 3, 3, 4),
			[][]int{
				{1},
				{2, 2},
				{3, 3, 3},
				{4},
			},
		},
		{nil, nil},
	}

	for _, tt := range tests {
		order := levelOrder(tt.tree)
		if reflect.DeepEqual(order, tt.order) == false {
			t.Errorf("levelOrder(%v) return %v, want %v", tt.tree, order, tt.order)
		}
	}
}
