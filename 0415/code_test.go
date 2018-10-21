package leetcode

import "testing"

func TestAddStrings(t *testing.T) {
	var tests = []struct {
		num1, num2 string
		sum        string
	}{
		{"0", "0", "0"},
		{"9", "9", "18"},
		{"9", "99", "108"},
		{"123456789", "987654321", "1111111110"},
	}

	for _, tt := range tests {
		sum := addStrings(tt.num1, tt.num2)
		if sum != tt.sum {
			t.Errorf("addStrings(%v, %v) return %v, want %v", tt.num1, tt.num2, sum, tt.sum)
		}
	}
}
