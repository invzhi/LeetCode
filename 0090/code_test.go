package leetcode

import (
	"reflect"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	var tests = []struct {
		nums []int
		sets [][]int
	}{
		{[]int{1, 2, 2}, [][]int{nil, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}},
	}

	for _, tt := range tests {
		sets := subsetsWithDup(tt.nums)
		if reflect.DeepEqual(sets, tt.sets) == false {
			t.Errorf("subsetsWithDup(%v) return %v, want %v", tt.nums, sets, tt.sets)
		}
	}
}
