package leetcode

import (
	"reflect"
	"testing"
)

func TestGrayCode(t *testing.T) {
	var tests = []struct {
		n     int
		codes []int
	}{
		{0, []int{0}},
		{1, []int{0, 1}},
		{2, []int{0, 1, 3, 2}},
		{3, []int{0, 1, 3, 2, 6, 7, 5, 4}},
	}

	for _, tt := range tests {
		codes := grayCode(tt.n)
		if reflect.DeepEqual(codes, tt.codes) == false {
			t.Errorf("grayCode(%v) return %v, want %v", tt.n, codes, tt.codes)
		}
	}
}
