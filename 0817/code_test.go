package leetcode

import "testing"

func TestNumComponents(t *testing.T) {
	var tests = []struct {
		head *ListNode
		G    []int
		num  int
	}{
		{newList(0, 1, 2, 3), []int{0, 1, 3}, 2},
		{newList(0, 1, 2, 3, 4), []int{0, 3, 1, 4}, 2},
	}

	for _, tt := range tests {
		num := numComponents(tt.head, tt.G)
		if num != tt.num {
			t.Errorf("numComponents(%v, %v) return %v, want %v", tt.head, tt.G, num, tt.num)
		}
	}
}
