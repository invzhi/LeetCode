// 57. Insert Interval

// Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
// You may assume that the intervals were initially sorted according to their start times.

// Example 1:
// Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
// Output: [[1,5],[6,9]]

// Example 2:
// Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// Output: [[1,2],[3,10],[12,16]]
// Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

package leetcode

// Interval represents interval with two values.
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	for i, interval := range intervals {
		if newInterval.Start <= interval.End {
			if newInterval.End < interval.Start {
				return append(append(intervals[:i:i], newInterval), intervals[i:]...)
			}

			if newInterval.Start < interval.Start {
				intervals[i].Start = newInterval.Start
			}
			if newInterval.End > interval.End {
				intervals[i].End = newInterval.End
			}

			j := i + 1
			for j < len(intervals) && intervals[j].Start <= newInterval.End {
				j++
			}
			if intervals[j-1].End > intervals[i].End {
				intervals[i].End = intervals[j-1].End
			}
			return append(intervals[:i+1], intervals[j:]...)
		}
	}
	return append(intervals, newInterval)
}
