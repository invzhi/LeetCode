package leetcode

import "testing"

func TestLRUCache(t *testing.T) {
	var tests = []struct {
		put        bool
		key, value int
	}{
		{true, 1, 1},
		{true, 2, 2},
		{false, 1, 1},
		{true, 3, 2},
		{true, 3, 3},
		{false, 2, -1},
		{true, 4, 4},
		{false, 1, -1},
		{false, 3, 3},
		{false, 4, 4},
	}

	cache := Constructor(2)
	for _, tt := range tests {
		if tt.put {
			cache.Put(tt.key, tt.value)
			t.Logf("cache.Put(%v, %v)\n", tt.key, tt.value)
		} else {
			value := cache.Get(tt.key)
			t.Logf("cache.Get(%v) return %v\n", tt.key, value)
			if value != tt.value {
				t.Errorf("want %v", tt.value)
			}
		}
	}
}
