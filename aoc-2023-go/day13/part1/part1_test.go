package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"strings"
	"testing"
)

func TestInput(t *testing.T) {
	patterns := strings.Split(Input, "\n\n")
	horizontal, vertical := 0, 0
	for _, pattern := range patterns {
		h, v := findSymmetryLines(pattern)
		horizontal += h
		vertical += v
	}

	sum := horizontal*100 + vertical

	fmt.Println(sum)
}

func findSymmetryLines(pattern string) (horizontal, vertical int) {
	grid := util.ToGrid(pattern)
	horizontal = 0
	vertical = 0
	for row := 0; row < len(grid)-1; row++ {
		symmetry := true
		for i := 0; i <= (util.Min(row, len(grid)-row-2)); i++ {
			if !util.StringArrayEqual(grid[row-i], grid[row+i+1]) {
				symmetry = false
				break
			}
		}
		if symmetry {
			horizontal = row + 1
		}
	}

	flippedGrid := util.RotateRight(grid)
	for row := 0; row < len(flippedGrid)-1; row++ {
		symmetry := true
		for i := 0; i <= (util.Min(row, len(flippedGrid)-row-2)); i++ {
			if !util.StringArrayEqual(flippedGrid[row-i], flippedGrid[row+i+1]) {
				symmetry = false
				break
			}
		}
		if symmetry {
			vertical = row + 1
		}
	}

	//fmt.Printf("%s \nhorizontal: %d, vertical: %d\n", pattern, horizontal, vertical)
	return horizontal, vertical
}
