package util

import "strings"

func InGrid(row int, column int, grid [][]string) bool {
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

type Point struct {
	Row    int
	Column int
}

func Find(value string, grid [][]string) Point {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if grid[row][column] == value {
				return Point{Row: row, Column: column}
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
		for column := point.Column - 1; column < point.Column+1; column++ {
			if row == point.Row && column == point.Column {
				continue
			}
			if row < 0 || row > maxRow {
				continue
			}
			if column < 0 || column > maxColumn {
				continue
			}
			adjacentPoints = append(adjacentPoints, Point{Row: row, Column: column})
		}
	}

	return adjacentPoints
}
