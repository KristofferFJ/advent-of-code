package util

import (
	"math"
)

func intersects(A, B Point, P [2]float64) bool {
	if A.Column > B.Column {
		return intersects(B, A, P)
	}

	if P[1] == float64(A.Column) || P[1] == float64(B.Column) {
		P[1] += 0.0001
	}

	if P[1] > float64(B.Column) || P[1] < float64(A.Column) || P[0] >= math.Max(float64(A.Row), float64(B.Row)) {
		return false
	}

	if P[0] < math.Min(float64(A.Row), float64(B.Row)) {
		return true
	}

	red := (P[1] - float64(A.Column)) / (P[0] - float64(A.Row))
	blue := (float64(B.Column) - float64(A.Column)) / (float64(B.Row) - float64(A.Row))
	return red >= blue
}

func FigureContains(shape []Point, point Point) bool {
	inside := false
	for i := 0; i < len(shape); i++ {
		if intersects(shape[i], shape[(i+1)%len(shape)], [2]float64{float64(point.Row), float64(point.Column)}) {
			inside = !inside
		}
	}
	return inside
}
