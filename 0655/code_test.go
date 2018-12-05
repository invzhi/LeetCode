package leetcode

import (
	"reflect"
	"testing"
)

func TestPrintTree(t *testing.T) {
	var tests = []struct {
		tree  *TreeNode
		array [][]string
	}{
		{
			newTree(1, 2),
			[][]string{
				{"", "1", ""},
				{"2", "", ""},
			},
		},
		{
			newTree(1, 2, 3, nil, 4),
			[][]string{
				{"", "", "", "1", "", "", ""},
				{"", "2", "", "", "", "3", ""},
				{"", "", "4", "", "", "", ""},
			},
		},
		{
			newTree(5, 3, 6, 2, 4, nil, 7),
			[][]string{
				{"", "", "", "5", "", "", ""},
				{"", "3", "", "", "", "6", ""},
				{"2", "", "4", "", "", "", "7"},
			},
		},
		{
			newTree(1, 2, 5, 3, nil, nil, nil, 4),
			[][]string{
				{"", "", "", "", "", "", "", "1", "", "", "", "", "", "", ""},
				{"", "", "", "2", "", "", "", "", "", "", "", "5", "", "", ""},
				{"", "3", "", "", "", "", "", "", "", "", "", "", "", "", ""},
				{"4", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			},
		},
	}

	for _, tt := range tests {
		array := printTree(tt.tree)
		if reflect.DeepEqual(array, tt.array) == false {
			t.Errorf("printTree(%v) return %q, want %q", tt.tree, array, tt.array)
		}
	}
}
