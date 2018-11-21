package leetcode

import "testing"

func TestIsHappy(t *testing.T) {
	var tests = []struct {
		n     int
		happy bool
	}{
		{1, true},
		{19, true},
		{313, true},
		{314, false},
	}

	for _, tt := range tests {
		happy := isHappy(tt.n)
		if happy != tt.happy {
			t.Errorf("isHappy(%v) return %v, want %v", tt.n, happy, tt.happy)
		}
	}
}
