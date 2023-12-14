package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"testing"
)

func TestInput(t *testing.T) {
	grid := util.ToGrid(Input)
	results := make(map[int][]int)
	for i := 1; i <= 1000000000; i++ {
		grid = cycle(grid)
		fmt.Println(i, count(grid))
		results[count(grid)] = append(results[count(grid)], i)
	}
}

func cycle(grid [][]string) [][]string {
	//north
	north := util.RotateRight(grid)
	for _, row := range north {
		blockage := len(row)
		for i := len(row) - 1; i >= 0; i-- {
			if row[i] == "O" {
				row[i] = "."
				row[blockage-1] = "O"
				blockage = blockage - 1
			} else if row[i] == "#" {
				blockage = i
			}
		}
	}

	//west
	west := util.RotateRight(north)
	for _, row := range west {
		blockage := len(row)
		for i := len(row) - 1; i >= 0; i-- {
			if row[i] == "O" {
				row[i] = "."
				row[blockage-1] = "O"
				blockage = blockage - 1
			} else if row[i] == "#" {
				blockage = i
			}
		}
	}

	//south
	south := util.RotateRight(west)
	for _, row := range south {
		blockage := len(row)
		for i := len(row) - 1; i >= 0; i-- {
			if row[i] == "O" {
				row[i] = "."
				row[blockage-1] = "O"
				blockage = blockage - 1
			} else if row[i] == "#" {
				blockage = i
			}
		}
	}

	//east
	east := util.RotateRight(south)
	for _, row := range east {
		blockage := len(row)
		for i := len(row) - 1; i >= 0; i-- {
			if row[i] == "O" {
				row[i] = "."
				row[blockage-1] = "O"
				blockage = blockage - 1
			} else if row[i] == "#" {
				blockage = i
			}
		}
	}

	return east
}

func count(eastGrid [][]string) int {
	sum := 0
	northGrid := util.Duplicate(eastGrid)
	for _, row := range util.RotateRight(northGrid) {
		for index, cell := range row {
			if cell == "O" {
				sum += index + 1
			}
		}
	}
	return sum
}
