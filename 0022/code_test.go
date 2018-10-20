package leetcode

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	var tests = []struct {
		n     int
		pairs []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, tt := range tests {
		pairs := generateParenthesis(tt.n)
		if reflect.DeepEqual(pairs, tt.pairs) == false {
			t.Errorf("generateParenthesis(%v) return %v, want %v", tt.n, pairs, tt.pairs)
		}
	}
}
