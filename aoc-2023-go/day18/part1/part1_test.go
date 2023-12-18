package part1

import (
	"fmt"
	. "io.kristofferfj.github/aoc-2023-go/util"
	"strconv"
	"strings"
	"testing"
)

type Instruction struct {
	dir  Point
	dist int
}

func TestInput(t *testing.T) {
	instructions := parseInput(Input)
	minRows, minCols, maxRows, maxCols := determineDimensions(instructions)
	grid := make([][]string, maxRows-minRows+1)
	for i := range grid {
		grid[i] = make([]string, maxCols-minCols+1)
	}
	pos := Point{Row: -minRows, Col: -minCols}
	grid[pos.Row][pos.Col] = "#"
	var Vertices []Point
	for _, instruction := range instructions {
		Vertices = append(Vertices, pos)
		for i := 0; i < instruction.dist; i++ {
			pos.Row += instruction.dir.Row
			pos.Col += instruction.dir.Col
			grid[pos.Row][pos.Col] = "#"
		}
	}

	var figure []Point
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == "#" {
				figure = append(figure, Point{Row: row, Col: col})
			}
		}
	}

	sum := 0
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if PointInList(Point{Row: row, Col: column}, figure) {
				continue
			}
			if !FigureContains(Vertices, Point{Row: row, Col: column}) {
				sum++
			}
		}
	}

	fmt.Println(len(grid)*len(grid[0]) - sum)
}

func parseInput(input string) []Instruction {
	var instructions []Instruction
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")
		dist, _ := strconv.Atoi(split[1])
		dir := Point{}
		switch split[0] {
		case "U":
			dir = UP
		case "D":
			dir = DOWN
		case "L":
			dir = LEFT
		case "R":
			dir = RIGHT
		}
		instructions = append(instructions, Instruction{dir, dist})
	}
	return instructions
}

func determineDimensions(instructions []Instruction) (minRow, minCol, maxRow, maxCol int) {
	rows := 0
	cols := 0
	for _, instruction := range instructions {
		if instruction.dir.Row != 0 {
			rows += instruction.dist * instruction.dir.Row
			if rows > maxRow {
				maxRow = rows
			}
			if rows < minRow {
				minRow = rows
			}
		} else {
			cols += instruction.dist * instruction.dir.Col
			if cols > maxCol {
				maxCol = cols
			}
			if cols < minCol {
				minCol = cols
			}
		}
	}
	return minRow, minCol, maxRow, maxCol
}
