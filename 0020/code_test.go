package leetcode

import "testing"

func TestIsValid(t *testing.T) {
	var tests = []struct {
		s     string
		valid bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, tt := range tests {
		valid := isValid(tt.s)
		if valid != tt.valid {
			t.Errorf("isValid(%v) return %v, want %v", tt.s, valid, tt.valid)
		}
	}
}
