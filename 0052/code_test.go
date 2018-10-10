package leetcode

import (
	"reflect"
	"testing"
)

func TestTotalNQueens(t *testing.T) {
	var tests = []struct {
		n   int
		cnt int
	}{
		{4, 2},
	}

	for _, tt := range tests {
		cnt := totalNQueens(tt.n)
		if reflect.DeepEqual(cnt, tt.cnt) == false {
			t.Errorf("totalNQueens(%v) return %v, want %v", tt.n, cnt, tt.cnt)
		}
	}
}
