package leetcode

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	var tests = []struct {
		before []string
		after  []string
	}{
		{[]string{"Hello", "Alaska", "Dad", "Peace"}, []string{"Alaska", "Dad"}},
	}

	for _, tt := range tests {
		before := make([]string, len(tt.before))
		copy(before, tt.before)

		after := findWords(tt.before)
		if reflect.DeepEqual(after, tt.after) == false {
			t.Errorf("findWords(%v) return %v, want %v", before, after, tt.after)
		}
	}
}
