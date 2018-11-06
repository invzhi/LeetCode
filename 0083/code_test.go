package leetcode

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	var tests = []struct {
		before []int
		after  *ListNode
	}{
		{[]int{1, 1, 2}, newList(1, 2)},
		{[]int{1, 1, 2, 3, 3}, newList(1, 2, 3)},
	}

	for _, tt := range tests {
		after := deleteDuplicates(newList(tt.before...))
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("deleteDuplicates(%v) return %v, want %v", newList(tt.before...), after, tt.after)
		}
	}
}
