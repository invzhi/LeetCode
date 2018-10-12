package leetcode

import (
	"reflect"
	"testing"
)

func TestLetterCasePermutation(t *testing.T) {
	var tests = []struct {
		S string
		s []string
	}{
		{"a1b2", []string{"a1b2", "A1b2", "a1B2", "A1B2"}},
		{"3z4", []string{"3z4", "3Z4"}},
		{"3Z4", []string{"3Z4", "3z4"}},
		{"12345", []string{"12345"}},
	}

	for _, tt := range tests {
		s := letterCasePermutation(tt.S)
		if reflect.DeepEqual(s, tt.s) == false {
			t.Errorf("letterCasePermutation(%v) return %v, want %v", tt.S, s, tt.s)
		}
	}
}
