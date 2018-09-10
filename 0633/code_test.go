package leetcode

import "testing"

func TestJudgeSquareSum(t *testing.T) {
	var tests = []struct {
		c           int
		isSquareSum bool
	}{
		{3, false},
		{5, true},
		{8, true},
		{21, false},
		{19801, true},
		{23050, true},
	}

	for _, tt := range tests {
		isSquareSum := judgeSquareSum(tt.c)
		if isSquareSum != tt.isSquareSum {
			t.Errorf("judgeSquareSum(%v) return %v, want %v", tt.c, isSquareSum, tt.isSquareSum)
		}
	}
}
