package leetcode

import (
	"math"
	"testing"
)

func TestRandomIndex(t *testing.T) {
	const times = 10000

	var tests = []struct {
		nums []int
	}{
		{[]int{1, 2, 3, 3, 3}},
		{[]int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 5, 5, 5, 5}},
		{[]int{1, 2, 3, 1, 2, 3, 1, 2, 3}},
	}

	for _, tt := range tests {
		cnt := make(map[int]int)
		for _, num := range tt.nums {
			cnt[num]++
		}

		s := Constructor(tt.nums)

		for target, c := range cnt {
			indexs := make(map[int]int)
			for i := 0; i < times; i++ {
				indexs[s.Pick(target)]++
			}

			want := 1 / float64(c)
			for i, n := range indexs {
				got := float64(n) / times
				if math.Abs(got-want) > 1e-2 {
					t.Errorf("%v: the probability of index %v is %v, want %v", tt.nums, i, got, want)
				}
			}
		}
	}
}
