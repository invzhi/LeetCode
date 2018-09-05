package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	var tests = []struct {
		roman string
		num   int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tt := range tests {
		num := romanToInt(tt.roman)
		if num != tt.num {
			t.Errorf("romanToInt(%v) return %v, want %v", tt.roman, num, tt.num)
		}
	}
}
