package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"testing"
)

func TestInput(t *testing.T) {
	sum := 0
	flippedGrid := util.RotateRight(util.ToGrid(Input))
	for _, row := range flippedGrid {
		blockage := len(row)
		for i := len(row) - 1; i >= 0; i-- {
			if row[i] == "O" {
				sum += blockage
				blockage = blockage - 1
			}
			if row[i] == "#" {
				blockage = i
			}
		}
	}
	fmt.Println(sum)
}
