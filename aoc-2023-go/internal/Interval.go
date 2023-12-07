package internal

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
	return intervals
}

func (interval Interval) RemoveSlice(other Interval) []Interval {
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

	return []Interval{interval}
}
