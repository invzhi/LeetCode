package leetcode

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	var tests = []struct {
		vs []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{1, 2, 2}},
		{[]interface{}{1, nil, 2, 3}},
		{[]interface{}{1, 2, 2, nil, nil, 3, 3}},
		{[]interface{}{1, 2, 2, 3, nil, nil, 3, 4}},
		{[]interface{}{1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6, nil, 7, nil, 8, nil, 9}},
	}

	for _, tt := range tests {
		s := newTree(tt.vs...).String()
		if s != fmt.Sprint(tt.vs) {
			t.Errorf("newTree(%v) return %v", tt.vs, s)
		}
	}
}
