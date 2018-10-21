package leetcode

import "testing"

func TestConvertToTitle(t *testing.T) {
	var tests = []struct {
		number int
		title  string
	}{
		{1, "A"},
		{28, "AB"},
		{701, "ZY"},
	}

	for _, tt := range tests {
		title := convertToTitle(tt.number)
		if title != tt.title {
			t.Errorf("convertToTitle(%v) return %v, want %v", tt.number, title, tt.title)
		}
	}
}
