package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"testing"
)

func TestInput(t *testing.T) {
	grid := util.ToGrid(Input)
	horizontallyAlteredGrid := make([][]string, len(grid))
	copy(horizontallyAlteredGrid, grid)
	addedLines := 0
	for row, line := range grid {
		if emptyLine(line) {
			horizontallyAlteredGrid = util.Insert(horizontallyAlteredGrid, row+addedLines, line)
			addedLines++
		}
	}
	horizontallyAlteredGrid = util.RotateRight(horizontallyAlteredGrid)
	verticallyAlteredGrid := make([][]string, len(horizontallyAlteredGrid))
	copy(verticallyAlteredGrid, horizontallyAlteredGrid)
	addedLines = 0
	for row, line := range horizontallyAlteredGrid {
		if emptyLine(line) {
			verticallyAlteredGrid = util.Insert(verticallyAlteredGrid, row+addedLines, line)
			addedLines++
		}
	}
	verticallyAlteredGrid = util.RotateLeft(verticallyAlteredGrid)

	var galaxies [][2]int
	for row := 0; row < len(verticallyAlteredGrid); row++ {
		for column := 0; column < len(verticallyAlteredGrid[row]); column++ {
			if verticallyAlteredGrid[row][column] == string('#') {
				galaxies = append(galaxies, [2]int{row, column})
			}
		}
	}

	sum := 0
	for from := 0; from < len(galaxies)-1; from++ {
		for to := from + 1; to < len(galaxies); to++ {
			distance := util.Abs(galaxies[from][0]-galaxies[to][0]) + util.Abs(galaxies[from][1]-galaxies[to][1])
			sum += distance
		}
	}
	fmt.Println(sum)
}

func emptyLine(line []string) bool {
	for _, elem := range line {
		if elem != "." {
			return false
		}
	}
	return true
}
