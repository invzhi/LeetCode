package leetcode

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	var tests = []struct {
		digits       string
		combinations []string
	}{
		{"", nil},
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}

	for _, tt := range tests {
		combinations := letterCombinations(tt.digits)
		if reflect.DeepEqual(combinations, tt.combinations) == false {
			t.Errorf("letterCombinations(%v) return %v, want %v", tt.digits, combinations, tt.combinations)
		}
	}
}
