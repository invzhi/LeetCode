package leetcode

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	var tests = []struct {
		rowIndex int
		row      []int
	}{
		{3, []int{1, 3, 3, 1}},
	}

	for _, tt := range tests {
		row := getRow(tt.rowIndex)
		if reflect.DeepEqual(row, tt.row) == false {
			t.Errorf("getRow(%v) return %v, want %v", tt.rowIndex, row, tt.row)
		}
	}
}
