package util

import (
	"strconv"
	"strings"
)

func InGrid[T any](row int, column int, grid [][]T) bool {
	return row >= 0 && row < len(grid) && column >= 0 && column < len(grid[0])
}

func ToGrid(s string) [][]string {
	lines := strings.Split(s, "\n")
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func ToIntGrid(s string) [][]int {
	lines := strings.Split(s, "\n")
	var grid [][]int
	for _, line := range lines {
		var row []int
		for _, char := range strings.Split(line, "") {
			num, _ := strconv.Atoi(char)
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	return grid
}

type Point struct {
	Row int
	Col int
}

func Find(value string, grid [][]string) Point {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if grid[row][column] == value {
				return Point{Row: row, Col: column}
			}
		}
	}
	panic("not found")
}

func Adjacent(point Point, grid [][]string) []Point {
	maxColumn := len(grid[0]) - 1
	maxRow := len(grid) - 1
	var adjacentPoints []Point
	for row := point.Row - 1; row < point.Row+1; row++ {
		for column := point.Col - 1; column < point.Col+1; column++ {
			if row == point.Row && column == point.Col {
				continue
			}
			if row < 0 || row > maxRow {
				continue
			}
			if column < 0 || column > maxColumn {
				continue
			}
			adjacentPoints = append(adjacentPoints, Point{Row: row, Col: column})
		}
	}

	return adjacentPoints
}

var (
	UP    = Point{Row: -1}
	DOWN  = Point{Row: 1}
	LEFT  = Point{Col: -1}
	RIGHT = Point{Col: 1}
)
