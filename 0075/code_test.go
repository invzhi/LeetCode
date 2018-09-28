package leetcode

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	var tests = []struct {
		colors       []int
		sortedColors []int
	}{
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}

	for _, tt := range tests {
		colors := make([]int, len(tt.colors))
		copy(colors, tt.colors)

		sortColors(tt.colors)
		if reflect.DeepEqual(tt.colors, tt.sortedColors) == false {
			t.Errorf("sortColors(%v) return %v, want %v", colors, tt.colors, tt.sortedColors)
		}
	}
}
