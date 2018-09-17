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

func TestIncreasingBST(t *testing.T) {
	var tests = []struct {
		vals []interface{}
		tree *TreeNode
	}{
		{[]interface{}{3, 2, 4, 1}, newTree(1, nil, 2, nil, nil, nil, 3, nil, nil, nil, nil, nil, nil, nil, 4)},
	}

	for _, tt := range tests {
		tree := increasingBST(newTree(tt.vals...))
		if reflect.DeepEqual(tree, tt.tree) == false {
			t.Errorf("increasingBST(%v) return %v, want %v", newTree(tt.vals...), tree, tt.tree)
		}
	}
}
