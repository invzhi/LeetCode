// 56. Merge Intervals

// Given a collection of intervals, merge all overlapping intervals.

// Example 1:
// Input: [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

// Example 2:
// Input: [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considerred overlapping.

package leetcode

import "sort"

// Interval represents interval with two values.
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })

	var i int
	merged := []Interval{intervals[0]}
	for _, interval := range intervals {
		if merged[i].End < interval.Start {
			merged = append(merged, interval)
			i++
		} else if interval.End > merged[i].End {
			merged[i].End = interval.End
		}
	}
	return merged
}
