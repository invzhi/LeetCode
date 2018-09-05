package leetcode

import "testing"

func TestIntToRoman(t *testing.T) {
	var tests = []struct {
		num   int
		roman string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, tt := range tests {
		roman := intToRoman(tt.num)
		if roman != tt.roman {
			t.Errorf("intToroman(%v) return %v, want %v", tt.num, roman, tt.roman)
		}
	}
}
