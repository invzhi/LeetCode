package leetcode

import "testing"

func TestIsOneBitCharacter(t *testing.T) {
	var tests = []struct {
		bits     []int
		isOneBit bool
	}{
		{[]int{1, 0, 0}, true},
		{[]int{1, 1, 1, 0}, false},
	}

	for _, tt := range tests {
		isOneBit := isOneBitCharacter(tt.bits)
		if isOneBit != tt.isOneBit {
			t.Errorf("isOneBitCharacter(%v) return %v, want %v", tt.bits, isOneBit, tt.isOneBit)
		}
	}
}
