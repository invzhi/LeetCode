package leetcode

import "testing"

func TestToHex(t *testing.T) {
	var tests = []struct {
		num int
		hex string
	}{
		{26, "1a"},
		{16, "10"},
		{-1, "ffffffff"},
		{0, "0"},
	}

	for _, tt := range tests {
		hex := toHex(tt.num)
		if hex != tt.hex {
			t.Errorf("toHex(%v) return %v, want %v", tt.num, hex, tt.hex)
		}
	}
}
