package internal

type Interval struct {
	startInclusive int
	endInclusive   int
}

func (this Interval) intersection(other Interval) Interval {
	if other.endInclusive < this.startInclusive || this.endInclusive < other.startInclusive {
		return Interval{}
	}
	return Interval{
		startInclusive: Max(this.startInclusive, other.startInclusive),
		endInclusive:   Min(this.endInclusive, other.endInclusive),
	}
}
