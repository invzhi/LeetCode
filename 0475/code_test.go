package leetcode

import "testing"

func TestFindRadius(t *testing.T) {
	var tests = []struct {
		houses, heaters []int
		radius          int
	}{
		{[]int{1, 5}, []int{2}, 3},
		{[]int{1, 5}, []int{10}, 9},
		{[]int{1, 2, 3}, []int{2}, 1},
		{[]int{1, 2, 3, 4}, []int{1, 4}, 1},
	}

	for _, tt := range tests {
		radius := findRadius(tt.houses, tt.heaters)
		if radius != tt.radius {
			t.Errorf("findRadius(%v, %v) return %v, want %v", tt.houses, tt.heaters, radius, tt.radius)
		}
	}
}
