package leetcode

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	var tests = []struct {
		nums []int
		sets [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{}, {3}, {2}, {2, 3}, {1}, {1, 3}, {1, 2}, {1, 2, 3}}},
		{[]int{9, 0, 3, 5, 7}, [][]int{{}, {7}, {5}, {5, 7}, {3}, {3, 7}, {3, 5}, {3, 5, 7}, {0}, {0, 7}, {0, 5}, {0, 5, 7}, {0, 3}, {0, 3, 7}, {0, 3, 5}, {0, 3, 5, 7}, {9}, {9, 7}, {9, 5}, {9, 5, 7}, {9, 3}, {9, 3, 7}, {9, 3, 5}, {9, 3, 5, 7}, {9, 0}, {9, 0, 7}, {9, 0, 5}, {9, 0, 5, 7}, {9, 0, 3}, {9, 0, 3, 7}, {9, 0, 3, 5}, {9, 0, 3, 5, 7}}},
	}

	for _, tt := range tests {
		sets := subsets(tt.nums)
		if reflect.DeepEqual(sets, tt.sets) == false {
			t.Errorf("subsets(%v) return %v, want %v", tt.nums, sets, tt.sets)
		}
	}
}
