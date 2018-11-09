package leetcode

import "testing"

func TestKthLargest(t *testing.T) {
	type add struct{ val, kth int }

	var tests = []struct {
		k    int
		nums []int
		adds []add
	}{
		{
			3, []int{4, 5, 8, 2},
			[]add{
				{3, 4},
				{5, 5},
				{10, 5},
				{9, 8},
				{4, 8},
			},
		},
		{
			1, []int{},
			[]add{
				{-3, -3},
				{-2, -2},
				{-4, -2},
				{0, 0},
				{4, 4},
			},
		},
		{
			3, []int{3, 2},
			[]add{
				{1, 1},
				{5, 2},
				{6, 3},
				{4, 4},
			},
		},
	}

	for _, tt := range tests {
		kl := Constructor(tt.k, tt.nums)
		t.Logf("Constructor(%v, %v)", tt.k, tt.nums)
		for _, add := range tt.adds {
			kth := kl.Add(add.val)
			t.Logf("KthLargest.Add(%v) return %v", add.val, kth)
			if kth != add.kth {
				t.Errorf("want %v", add.kth)
			}
		}
	}
}
