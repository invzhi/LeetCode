package leetcode

import (
	"math"
	"testing"
)

func TestRandomNode(t *testing.T) {
	const times = 10000

	var tests = []struct {
		vals []int
	}{
		{[]int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		cnt := make(map[int]int, len(tt.vals))

		s := Constructor(newList(tt.vals...))
		for i := 0; i < times; i++ {
			cnt[s.GetRandom()]++
		}

		want := 1 / float64(len(tt.vals))
		for v, n := range cnt {
			got := float64(n) / times
			if math.Abs(got-want) > 1e-2 {
				t.Errorf("%v: the probability of %v is %v, want %v", tt.vals, v, got, want)
			}
		}
	}
}
