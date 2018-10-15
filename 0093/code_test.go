package leetcode

import (
	"reflect"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	var tests = []struct {
		s   string
		IPs []string
	}{
		{"0000", []string{"0.0.0.0"}},
		{"0111", []string{"0.1.1.1"}},
		{"11101", []string{"1.1.10.1", "1.11.0.1", "11.1.0.1"}},
		{"11111", []string{"1.1.1.11", "1.1.11.1", "1.11.1.1", "11.1.1.1"}},
		{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
	}

	for _, tt := range tests {
		IPs := restoreIpAddresses(tt.s)
		if reflect.DeepEqual(IPs, tt.IPs) == false {
			t.Errorf("restoreIpAddresses(%v) return %v, want %v", tt.s, IPs, tt.IPs)
		}
	}
}
