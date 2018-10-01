package leetcode

import "testing"

func TestMiddleNode(t *testing.T) {
	var tests = []struct {
		list   *ListNode
		middle int
	}{
		{newList(1, 2, 3, 4, 5), 3},
		{newList(1, 2, 3, 4, 5, 6), 4},
	}

	for _, tt := range tests {
		middle := middleNode(tt.list)
		if middle.Val != tt.middle {
			t.Errorf("middleNode(%v) return %v, want %v", tt.list, middle.Val, tt.middle)
		}
	}
}
