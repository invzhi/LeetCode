package leetcode

import (
	"reflect"
	"testing"
)

type twoSumFunc struct {
	f    func([]int, int) []int
	name string
}

func testTwoSum(t *testing.T, f twoSumFunc) {
	var tests = []struct {
		numbers []int
		target  int
		indices []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 5, 6, 7}, 12, []int{2, 4}},
		{[]int{2, 6, 8, 23}, 16, nil},
	}

	for _, tt := range tests {
		indices := f.f(tt.numbers, tt.target)
		if reflect.DeepEqual(indices, tt.indices) == false {
			t.Errorf("%v(%v, %v) return %v, want %v", f.name, tt.numbers, tt.target, indices, tt.indices)
		}
	}
}

func TestTwoSum(t *testing.T) {
	testTwoSum(t, twoSumFunc{f: twoSum, name: "twoSum"})
}

func TestTwoSum1(t *testing.T) {
	testTwoSum(t, twoSumFunc{f: twoSum1, name: "twoSum"})
}
