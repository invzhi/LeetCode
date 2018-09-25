package leetcode

import "fmt"

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
	if v, ok := vs[0].(int); ok {
		nodes[0] = &TreeNode{Val: v}
	}
	for i, j := 0, 1; j < len(vs); i++ {
		if nodes[i] == nil {
			continue
		}

		if v, ok := vs[j].(int); ok {
			nodes[j] = &TreeNode{Val: v}
			nodes[i].Left = nodes[j]
		}

		if j+1 == len(vs) {
			break
		}
		if v, ok := vs[j+1].(int); ok {
			nodes[j+1] = &TreeNode{Val: v}
			nodes[i].Right = nodes[j+1]
		}

		j += 2
	}
	return nodes[0]
}
