package util

import (
	"math"
)

func intersects(A, B Point, P [2]float64) bool {
	if A.Col > B.Col {
		return intersects(B, A, P)
	}

	if P[1] == float64(A.Col) || P[1] == float64(B.Col) {
		P[1] += 0.0001
	}

	if P[1] > float64(B.Col) || P[1] < float64(A.Col) || P[0] >= math.Max(float64(A.Row), float64(B.Row)) {
		return false
	}

	if P[0] < math.Min(float64(A.Row), float64(B.Row)) {
		return true
	}

	red := (P[1] - float64(A.Col)) / (P[0] - float64(A.Row))
	blue := (float64(B.Col) - float64(A.Col)) / (float64(B.Row) - float64(A.Row))
	return red >= blue
}

func FigureContains(shape []Point, point Point) bool {
	inside := false
	for i := 0; i < len(shape); i++ {
		if intersects(shape[i], shape[(i+1)%len(shape)], [2]float64{float64(point.Row), float64(point.Col)}) {
			inside = !inside
		}
	}
	return inside
}

func CalculateArea(vertices []Point) (sum int) {
	for i := 0; i < len(vertices); i++ {
		sum += vertices[i].Row*vertices[(i+1)%len(vertices)].Col - vertices[i].Col*vertices[(i+1)%len(vertices)].Row
	}
	return Abs(sum / 2)
}

func CalculateEnclosedPoints(vertices []Point) (sum int) {
	for i := 0; i < len(vertices); i++ {
		sum += Abs(vertices[i].Row-vertices[(i+1)%len(vertices)].Row) + Abs(vertices[i].Col-vertices[(i+1)%len(vertices)].Col)
	}
	return CalculateArea(vertices) + (sum+2)/2
}
