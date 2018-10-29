package leetcode

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	var tests = []struct {
		list   []int
		swaped *ListNode
	}{
		{[]int{1}, newList(1)},
		{[]int{1, 2, 3, 4}, newList(2, 1, 4, 3)},
		{[]int{1, 2, 3, 4, 5}, newList(2, 1, 4, 3, 5)},
	}

	for _, tt := range tests {
		swaped := swapPairs(newList(tt.list...))
		if reflect.DeepEqual(swaped, tt.swaped) == false {
			t.Errorf("swapPairs(%v) return %v, want %v", newList(tt.list...), swaped, tt.swaped)
		}
	}
}
