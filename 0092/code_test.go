package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	var tests = []struct {
		before []int
		m, n   int
		after  *ListNode
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, newList(1, 4, 3, 2, 5)},
	}

	for _, tt := range tests {
		after := reverseBetween(newList(tt.before...), tt.m, tt.n)
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("reverseBetween(%v, %v, %v) return %v, want %v", newList(tt.before...), tt.m, tt.n, after, tt.after)
		}
	}
}
