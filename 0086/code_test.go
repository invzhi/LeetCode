package leetcode

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	var tests = []struct {
		before []int
		x      int
		after  *ListNode
	}{
		{[]int{1, 4, 3, 2, 5, 2}, 3, newList(1, 2, 2, 4, 3, 5)},
	}

	for _, tt := range tests {
		after := partition(newList(tt.before...), tt.x)
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("partition(%v, %v) return %v, want %v", newList(tt.before...), tt.x, after, tt.after)
		}
	}
}
