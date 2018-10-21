package leetcode

import "testing"

func TestTitleToNumber(t *testing.T) {
	var tests = []struct {
		title  string
		number int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
	}

	for _, tt := range tests {
		number := titleToNumber(tt.title)
		if number != tt.number {
			t.Errorf("titleToNumber(%v) return %v, want %v", tt.title, number, tt.number)
		}
	}
}
