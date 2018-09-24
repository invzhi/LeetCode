package leetcode

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	var tests = []struct {
		digits []int
		plus   []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9, 9, 9, 9}, []int{1, 0, 0, 0, 0}},
	}

	for _, tt := range tests {
		digits := make([]int, len(tt.digits))
		copy(digits, tt.digits)

		plus := plusOne(tt.digits)
		if reflect.DeepEqual(plus, tt.plus) == false {
			t.Errorf("plusOne(%v) return %v, want %v", digits, plus, tt.plus)
		}
	}
}
