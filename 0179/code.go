// 179. Largest Number

// Given a list of non negative integers, arrange them such that they form the largest number.

// Example 1:
// Input: [10,2]
// Output: "210"

// Example 2:
// Input: [3,30,34,5,9]
// Output: "9534330"

// Note: The result may be very large, so you need to return a string instead of an integer.

package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

type ByOrder []string

func (a ByOrder) Len() int {
	return len(a)
}

func (a ByOrder) Less(i, j int) bool {
	return a[j]+a[i] < a[i]+a[j]
}

func (a ByOrder) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}
	sort.Sort(ByOrder(ss))

	if ss[0] == "0" {
		return "0"
	}

	var b strings.Builder
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}
