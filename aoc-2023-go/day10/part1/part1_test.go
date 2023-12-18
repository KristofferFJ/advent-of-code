package part1

import (
	"fmt"
	. "io.kristofferfj.github/aoc-2023-go/util"
	"testing"
)

func TestInput(t *testing.T) {
	grid := ToGrid(Input)
	startingPoint := Find("S", grid)
	path := []Point{startingPoint}

	if leftValid(grid, startingPoint) {
		loop, _ := findLoop(grid, append(path, Point{Row: startingPoint.Row, Col: startingPoint.Col - 1}))
		fmt.Println(len(loop))
	}
	if rightValid(grid, startingPoint) {
		loop, _ := findLoop(grid, append(path, Point{Row: startingPoint.Row, Col: startingPoint.Col + 1}))
		fmt.Println(len(loop))
	}
	if upValid(grid, startingPoint) {
		loop, _ := findLoop(grid, append(path, Point{Row: startingPoint.Row - 1, Col: startingPoint.Col}))
		fmt.Println(len(loop))
	}
	if downValid(grid, startingPoint) {
		loop, _ := findLoop(grid, append(path, Point{Row: startingPoint.Row + 1, Col: startingPoint.Col}))
		fmt.Println(len(loop))
	}

	fmt.Println(path)
}

func TestTestInput(t *testing.T) {
	grid := ToGrid(InputTest)
	startingPoint := Find("S", grid)
	path := []Point{startingPoint}

	if leftValid(grid, startingPoint) {
		loop, _ := findLoop(grid, append(path, Point{Row: startingPoint.Row, Col: startingPoint.Col - 1}))
		fmt.Println(len(loop))
	}
	if rightValid(grid, startingPoint) {
		loop, _ := findLoop(grid, append(path, Point{Row: startingPoint.Row, Col: startingPoint.Col + 1}))
		fmt.Println(len(loop))
	}
	if upValid(grid, startingPoint) {
		loop, _ := findLoop(grid, append(path, Point{Row: startingPoint.Row - 1, Col: startingPoint.Col}))
		fmt.Println(len(loop))
	}
	if downValid(grid, startingPoint) {
		loop, _ := findLoop(grid, append(path, Point{Row: startingPoint.Row + 1, Col: startingPoint.Col}))
		fmt.Println(len(loop))
	}

	fmt.Println(path)
}

func findLoop(grid [][]string, path []Point) ([]Point, bool) {
	lastPoint := path[len(path)-1]
	if len(path) > 1 && grid[lastPoint.Row][lastPoint.Col] == "S" {
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
	switch lastValue := grid[lastPoint.Row][lastPoint.Col]; lastValue {
	case "-":
		if leftValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row, Col: lastPoint.Col - 1})
		}
		if rightValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row, Col: lastPoint.Col + 1})
		}
	case "|":
		if upValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row - 1, Col: lastPoint.Col})
		}
		if downValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row + 1, Col: lastPoint.Col})
		}
	case "F":
		if downValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row + 1, Col: lastPoint.Col})
		}
		if rightValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row, Col: lastPoint.Col + 1})
		}
	case "L":
		if rightValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row, Col: lastPoint.Col + 1})
		}
		if upValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row - 1, Col: lastPoint.Col})
		}
	case "J":
		if upValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row - 1, Col: lastPoint.Col})
		}
		if leftValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row, Col: lastPoint.Col - 1})
		}
	case "7":
		if leftValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row, Col: lastPoint.Col - 1})
		}
		if downValid(grid, lastPoint) {
			connected = append(connected, Point{Row: lastPoint.Row + 1, Col: lastPoint.Col})
		}
	}

	filtered := Filter(connected, func(point Point) bool {
		secondToLast := path[len(path)-2]
		if secondToLast.Row == point.Row && secondToLast.Col == point.Col {
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
	left := Point{Row: point.Row, Col: point.Col - 1}
	if !InGrid(left.Row, left.Col, grid) {
		return false
	}
	leftValue := grid[left.Row][left.Col]
	if leftValue == "-" || leftValue == "L" || leftValue == "F" || leftValue == "S" {
		return true
	}
	return false
}

func rightValid(grid [][]string, point Point) bool {
	right := Point{Row: point.Row, Col: point.Col + 1}
	if !InGrid(right.Row, right.Col, grid) {
		return false
	}
	rightValue := grid[right.Row][right.Col]
	if rightValue == "-" || rightValue == "J" || rightValue == "7" || rightValue == "S" {
		return true
	}
	return false
}

func upValid(grid [][]string, point Point) bool {
	up := Point{Row: point.Row - 1, Col: point.Col}
	if !InGrid(up.Row, up.Col, grid) {
		return false
	}
	upValue := grid[up.Row][up.Col]
	if upValue == "|" || upValue == "F" || upValue == "7" || upValue == "S" {
		return true
	}
	return false
}

func downValid(grid [][]string, point Point) bool {
	down := Point{Row: point.Row + 1, Col: point.Col}
	if !InGrid(down.Row, down.Col, grid) {
		return false
	}
	downValue := grid[down.Row][down.Col]
	if downValue == "|" || downValue == "L" || downValue == "J" || downValue == "S" {
		return true
	}
	return false
}
