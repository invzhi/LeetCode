package leetcode

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	var tests = []struct {
		image    [][]int
		sr, sc   int
		newColor int
		modified [][]int
	}{
		{
			[][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			1, 1, 2,
			[][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
	}

	for _, tt := range tests {
		image := make([][]int, len(tt.image))
		for i, r := range tt.image {
			image[i] = make([]int, len(r))
			copy(image[i], r)
		}

		modified := floodFill(tt.image, tt.sr, tt.sc, tt.newColor)
		if reflect.DeepEqual(modified, tt.modified) == false {
			t.Errorf("floodFill(%v, %v, %v, %v) return %v, want %v", image, tt.sr, tt.sc, tt.newColor, modified, tt.modified)
		}
	}
}
