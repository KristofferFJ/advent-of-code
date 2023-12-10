package part1

import (
	"fmt"
	. "io.kristofferfj.github/aoc-2023-go/util"
	"testing"
)

func TestInput(t *testing.T) {
	grid := ToGrid(Input)
	startingPoint := Find("S", grid)
	loop := findCorrectLoop(grid, startingPoint)
	var vertices []Point

	for i := 1; i < len(loop); i++ {
		if isVertex(i, loop) {
			vertices = append(vertices, loop[i])
		}
	}

	sum := 0
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if FigureContains(vertices, Point{Row: row, Column: column}) {
				if !PointInList(Point{Row: row, Column: column}, loop) {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}

func TestTestInput(t *testing.T) {
	grid := ToGrid(InputTest)
	startingPoint := Find("S", grid)
	loop := findCorrectLoop(grid, startingPoint)
	var vertices []Point

	for i := 1; i < len(loop); i++ {
		if isVertex(i, loop) {
			vertices = append(vertices, loop[i])
		}
	}

	sum := 0
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if FigureContains(vertices, Point{Row: row, Column: column}) {
				if !PointInList(Point{Row: row, Column: column}, loop) {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}

func isVertex(index int, loop []Point) bool {
	previous := loop[index].Row - loop[index-1].Row
	next := loop[(index+1)%len(loop)].Row - loop[index].Row

	if previous != next {
		return true
	}
	return false
}

func findCorrectLoop(grid [][]string, startingPoint Point) []Point {
	path := []Point{startingPoint}
	if leftValid(grid, startingPoint) {
		loop, valid := findLoop(grid, append(path, Point{Row: startingPoint.Row, Column: startingPoint.Column - 1}))
		if valid {
			return loop
		}
	}
	if rightValid(grid, startingPoint) {
		loop, valid := findLoop(grid, append(path, Point{Row: startingPoint.Row, Column: startingPoint.Column + 1}))
		if valid {
			return loop
		}
	}
	if upValid(grid, startingPoint) {
		loop, valid := findLoop(grid, append(path, Point{Row: startingPoint.Row - 1, Column: startingPoint.Column}))
		if valid {
			return loop
		}
	}
	if downValid(grid, startingPoint) {
		loop, valid := findLoop(grid, append(path, Point{Row: startingPoint.Row + 1, Column: startingPoint.Column}))
		if valid {
			return loop
		}
	}
	panic("not found")
}

func findLoop(grid [][]string, path []Point) ([]Point, bool) {
	lastPoint := path[len(path)-1]
	if len(path) > 1 && grid[lastPoint.Row][lastPoint.Column] == "S" {
		return path, true
	}

	next, exists := getNext(grid, path)
	if !exists {
		return []Point{}, false
	}

	return findLoop(grid, append(path, next))
}

func getNext(grid [][]string, path []Point) (Point, bool) {
	var connected []Point
	lastPoint := path[len(path)-1]
	switch lastValue := grid[lastPoint.Row][lastPoint.Column]; lastValue {
	case "-":
		if leftValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row, Column: lastPoint.Column - 1})
		}
		if rightValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row, Column: lastPoint.Column + 1})
		}
	case "|":
		if upValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row - 1, Column: lastPoint.Column})
		}
		if downValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row + 1, Column: lastPoint.Column})
		}
	case "F":
		if downValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row + 1, Column: lastPoint.Column})
		}
		if rightValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row, Column: lastPoint.Column + 1})
		}
	case "L":
		if rightValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row, Column: lastPoint.Column + 1})
		}
		if upValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row - 1, Column: lastPoint.Column})
		}
	case "J":
		if upValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row - 1, Column: lastPoint.Column})
		}
		if leftValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row, Column: lastPoint.Column - 1})
		}
	case "7":
		if leftValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row, Column: lastPoint.Column - 1})
		}
		if downValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row + 1, Column: lastPoint.Column})
		}
	}

	filtered := Filter(connected, func(point Point) bool {
		secondToLast := path[len(path)-2]
		if secondToLast.Row == point.Row && secondToLast.Column == point.Column {
			return false
		}
		return true
	})

	if len(filtered) == 0 {
		return Point{}, false
	}

	return filtered[0], true
}

func leftValid(grid [][]string, point Point) bool {
	left := Point{Row: point.Row, Column: point.Column - 1}
	if !InGrid(left.Row, left.Column, grid) {
		return false
	}
	leftValue := grid[left.Row][left.Column]
	if leftValue == "-" || leftValue == "L" || leftValue == "F" || leftValue == "S" {
		return true
	}
	return false
}

func rightValid(grid [][]string, point Point) bool {
	right := Point{Row: point.Row, Column: point.Column + 1}
	if !InGrid(right.Row, right.Column, grid) {
		return false
	}
	rightValue := grid[right.Row][right.Column]
	if rightValue == "-" || rightValue == "J" || rightValue == "7" || rightValue == "S" {
		return true
	}
	return false
}

func upValid(grid [][]string, point Point) bool {
	up := Point{Row: point.Row - 1, Column: point.Column}
	if !InGrid(up.Row, up.Column, grid) {
		return false
	}
	upValue := grid[up.Row][up.Column]
	if upValue == "|" || upValue == "F" || upValue == "7" || upValue == "S" {
		return true
	}
	return false
}

func downValid(grid [][]string, point Point) bool {
	down := Point{Row: point.Row + 1, Column: point.Column}
	if !InGrid(down.Row, down.Column, grid) {
		return false
	}
	downValue := grid[down.Row][down.Column]
	if downValue == "|" || downValue == "L" || downValue == "J" || downValue == "S" {
		return true
	}
	return false
}
