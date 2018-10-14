package leetcode

import "testing"

func TestToLowerCase(t *testing.T) {
	var tests = []struct {
		str string
		low string
	}{
		{"Hello", "hello"},
		{"here", "here"},
		{"LOVELY", "lovely"},
	}

	for _, tt := range tests {
		low := toLowerCase(tt.str)
		if low != tt.low {
			t.Errorf("toLowerCase(%v) return %v, want %v", tt.str, low, tt.low)
		}
	}
}
