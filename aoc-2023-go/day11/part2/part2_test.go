package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"testing"
)

func TestInput(t *testing.T) {
	grid := util.ToGrid(Input)
	var emptyRows []int
	var emptyColumns []int
	for row, line := range grid {
		if emptyLine(line) {
			emptyRows = append(emptyRows, row)
		}
	}
	for column, line := range util.RotateRight(grid) {
		if emptyLine(line) {
			emptyColumns = append(emptyColumns, column)
		}
	}

	var galaxies [][2]int
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			if grid[row][column] == string('#') {
				galaxies = append(galaxies, [2]int{row, column})
			}
		}
	}

	sum := 0
	for from := 0; from < len(galaxies)-1; from++ {
		for to := from + 1; to < len(galaxies); to++ {
			distance := calculateDistance(galaxies[from], galaxies[to], emptyRows, emptyColumns, 1000000-1)
			sum += distance
		}
	}
	fmt.Println(sum)
}

func calculateDistance(fromGalaxy [2]int, toGalaxy [2]int, emptyRows []int, emptyColumns []int, emptySpaces int) int {
	distance := util.Abs(fromGalaxy[0]-toGalaxy[0]) + util.Abs(fromGalaxy[1]-toGalaxy[1])
	emptyLinesTraversed := 0
	minRow := util.Min(fromGalaxy[0], toGalaxy[0])
	maxRow := util.Max(fromGalaxy[0], toGalaxy[0])
	minColumn := util.Min(fromGalaxy[1], toGalaxy[1])
	maxColumn := util.Max(fromGalaxy[1], toGalaxy[1])

	for _, emptyRow := range emptyRows {
		if emptyRow > minRow && emptyRow < maxRow {
			emptyLinesTraversed++
		}
	}
	for _, emptyColumn := range emptyColumns {
		if emptyColumn > minColumn && emptyColumn < maxColumn {
			emptyLinesTraversed++
		}
	}

	distance += emptyLinesTraversed * emptySpaces

	return distance
}

func emptyLine(line []string) bool {
	for _, elem := range line {
		if elem != "." {
			return false
		}
	}
	return true
}
