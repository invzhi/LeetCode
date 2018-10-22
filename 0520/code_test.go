package leetcode

import "testing"

func TestDetectCapitalUse(t *testing.T) {
	var tests = []struct {
		word  string
		right bool
	}{
		{"USA", true},
		{"leetcode", true},
		{"Google", true},
		{"FlaG", false},
	}

	for _, tt := range tests {
		right := detectCapitalUse(tt.word)
		if right != tt.right {
			t.Errorf("detectCapitalUse(%v) return %v, want %v", tt.word, right, tt.right)
		}
	}
}
