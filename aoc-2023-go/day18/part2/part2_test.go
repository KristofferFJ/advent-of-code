package part1

import (
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
	pos := Point{}
	edges := 0
	var Vertices = []Point{pos}
	for i := 0; i < len(instructions); i++ {
		pos.Row += instructions[i].dir.Row * instructions[i].dist
		pos.Col += instructions[i].dir.Col * instructions[i].dist
		edges += instructions[i].dist
		Vertices = append(Vertices, pos)
	}

	println(CalculateArea(Vertices) + (edges+2)/2)
}

func parseInput(input string) []Instruction {
	var instructions []Instruction
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")
		dist, _ := strconv.ParseInt(split[2][2:len(split[2])-2], 16, 32)
		dir := Point{}
		switch split[2][len(split[2])-2 : len(split[2])-1] {
		case "0":
			dir = RIGHT
		case "1":
			dir = DOWN
		case "2":
			dir = LEFT
		case "3":
			dir = UP
		}
		instructions = append(instructions, Instruction{dir, int(dist)})
	}
	return instructions
}
