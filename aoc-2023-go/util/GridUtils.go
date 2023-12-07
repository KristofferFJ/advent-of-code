package util

import "strings"

func InGrid(row int, column int, grid [][]string) bool {
	return row >= 0 && row < len(grid[0]) && column >= 0 && column < len(grid)
}

func ToGrid(s string) [][]string {
	lines := strings.Split(s, "\n")
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}
