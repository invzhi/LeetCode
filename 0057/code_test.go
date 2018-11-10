package leetcode

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	var tests = []struct {
		intervals    []Interval
		newIntervals Interval
		inserted     []Interval
	}{
		{[]Interval{{1, 5}}, Interval{0, 3}, []Interval{{0, 5}}},
		{[]Interval{{1, 5}}, Interval{0, 6}, []Interval{{0, 6}}},
		{[]Interval{{0, 5}, {9, 12}}, Interval{7, 16}, []Interval{{0, 5}, {7, 16}}},
		{[]Interval{{1, 3}, {6, 9}}, Interval{2, 5}, []Interval{{1, 5}, {6, 9}}},
		{[]Interval{{1, 3}, {6, 9}}, Interval{4, 5}, []Interval{{1, 3}, {4, 5}, {6, 9}}},
		{[]Interval{{1, 3}, {6, 9}}, Interval{-1, 0}, []Interval{{-1, 0}, {1, 3}, {6, 9}}},
		{[]Interval{{1, 3}, {6, 9}}, Interval{10, 11}, []Interval{{1, 3}, {6, 9}, {10, 11}}},
		{[]Interval{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, Interval{4, 8}, []Interval{{1, 2}, {3, 10}, {12, 16}}},
	}

	for _, tt := range tests {
		intervals := make([]Interval, len(tt.intervals))
		copy(intervals, tt.intervals)

		inserted := insert(tt.intervals, tt.newIntervals)
		if reflect.DeepEqual(inserted, tt.inserted) == false {
			t.Errorf("insert(%v, %v) return %v, want %v", intervals, tt.newIntervals, inserted, tt.inserted)
		}
	}
}
