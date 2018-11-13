package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	var tests = []struct {
		before []int
		n      int
		after  *ListNode
	}{
		{[]int{1, 2, 3, 4, 5}, 2, newList(1, 2, 3, 5)},
	}

	for _, tt := range tests {
		after := removeNthFromEnd(newList(tt.before...), tt.n)
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("removeNthFromEnd(%v) return %v, want %v", newList(tt.before...), after, tt.after)
		}
	}
}
