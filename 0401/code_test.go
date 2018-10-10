package leetcode

import (
	"reflect"
	"testing"
)

func TestReadBinaryWatch(t *testing.T) {
	var tests = []struct {
		num   int
		times []string
	}{
		{0, []string{"0:00"}},
		{1, []string{"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"}},
		{2, []string{"3:00", "5:00", "9:00", "1:01", "1:02", "1:04", "1:08", "1:16", "1:32", "6:00", "10:00", "2:01", "2:02", "2:04", "2:08", "2:16", "2:32", "4:01", "4:02", "4:04", "4:08", "4:16", "4:32", "8:01", "8:02", "8:04", "8:08", "8:16", "8:32", "0:03", "0:05", "0:09", "0:17", "0:33", "0:06", "0:10", "0:18", "0:34", "0:12", "0:20", "0:36", "0:24", "0:40", "0:48"}},
	}

	for _, tt := range tests {
		times := readBinaryWatch(tt.num)
		if reflect.DeepEqual(times, tt.times) == false {
			t.Errorf("readBinaryWatch(%v) return %v, want %v", tt.num, times, tt.times)
		}
	}
}
