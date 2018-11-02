package leetcode

import "testing"

func TestMultiply(t *testing.T) {
	var tests = []struct {
		num1, num2 string
		product    string
	}{
		{"0", "0", "0"},
		{"2", "3", "6"},
		{"0", "456", "0"},
		{"123", "456", "56088"},
	}

	for _, tt := range tests {
		product := multiply(tt.num1, tt.num2)
		if product != tt.product {
			t.Errorf("multiply(%v, %v) return %v, want %v", tt.num1, tt.num2, product, tt.product)
		}
	}
}
