package util

import "sort"

type Interval struct {
	StartInclusive int
	EndInclusive   int
}

func (interval Interval) Intersection(other Interval) (Interval, bool) {
	if other.EndInclusive < interval.StartInclusive || interval.EndInclusive < other.StartInclusive {
		return Interval{}, false
	}
	return Interval{
		StartInclusive: Max(interval.StartInclusive, other.StartInclusive),
		EndInclusive:   Min(interval.EndInclusive, other.EndInclusive),
	}, true
}

func RemoveSlices(intervals []Interval, slices []Interval) []Interval {
	modifiedIntervals := intervals
	for _, slice := range slices {
		modifiedIntervals = removeSliceFromIntervals(modifiedIntervals, slice)
	}
	return modifiedIntervals
}

func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].StartInclusive < intervals[j].StartInclusive
	})

	merged := []Interval{intervals[0]}

	for _, current := range intervals {
		lastMerged := &merged[len(merged)-1]

		if current.StartInclusive <= lastMerged.EndInclusive {
			if current.EndInclusive > lastMerged.EndInclusive {
				lastMerged.EndInclusive = current.EndInclusive
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

func removeSliceFromIntervals(intervals []Interval, slice Interval) []Interval {
	var modifiedIntervals []Interval
	for _, interval := range intervals {
		modifiedIntervals = append(modifiedIntervals, interval.removeSlice(slice)...)
	}
	return modifiedIntervals
}

func (interval Interval) removeSlice(other Interval) []Interval {
	if other.StartInclusive > interval.EndInclusive {
		return []Interval{interval}
	}

	if other.EndInclusive < interval.StartInclusive {
		return []Interval{interval}
	}

	if other.StartInclusive > interval.StartInclusive && other.EndInclusive < interval.EndInclusive {
		return []Interval{{
			StartInclusive: interval.StartInclusive,
			EndInclusive:   other.StartInclusive - 1,
		}, {
			StartInclusive: other.EndInclusive + 1,
			EndInclusive:   interval.EndInclusive,
		}}
	}

	if other.EndInclusive < interval.EndInclusive {
		return []Interval{{
			StartInclusive: other.EndInclusive + 1,
			EndInclusive:   interval.EndInclusive,
		}}
	}

	if other.StartInclusive > interval.StartInclusive {
		return []Interval{{
			StartInclusive: interval.StartInclusive,
			EndInclusive:   other.StartInclusive - 1,
		}}
	}

	return []Interval{}
}
