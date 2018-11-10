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
	if len(intervals) == 0 || newInterval.End < intervals[0].Start {
		return append([]Interval{newInterval}, intervals...)
	}
	for i := range intervals {
		if intervals[i].End < newInterval.Start && (i == len(intervals)-1 || newInterval.End < intervals[i+1].Start) {
			return append(append(intervals[:i+1:i+1], newInterval), intervals[i+1:]...)
		}
	}

	var i, j int
	for intervals[i].End < newInterval.Start {
		i++
	}
	if newInterval.Start < intervals[i].Start {
		intervals[i].Start = newInterval.Start
	}
	if newInterval.End > intervals[i].End {
		intervals[i].End = newInterval.End
	}
	for j = i + 1; j < len(intervals) && intervals[j].Start <= newInterval.End; j++ {
		if intervals[j].End > intervals[i].End {
			intervals[i].End = intervals[j].End
		}
	}
	return append(intervals[:i+1], intervals[j:]...)
}
