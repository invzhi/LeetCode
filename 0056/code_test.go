package leetcode

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	var tests = []struct {
		before []Interval
		after  []Interval
	}{
		{[]Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, []Interval{{1, 6}, {8, 10}, {15, 18}}},
		{[]Interval{{1, 4}, {4, 5}}, []Interval{{1, 5}}},
		{nil, nil},
	}

	for _, tt := range tests {
		before := make([]Interval, len(tt.before))
		copy(before, tt.before)

		after := merge(tt.before)
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("merge(%v) return %v, want %v", before, after, tt.after)
		}
	}
}
