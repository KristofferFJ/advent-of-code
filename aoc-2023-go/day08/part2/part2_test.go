package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"strings"
	"testing"
)

type Path struct {
	left  string
	right string
}

func parseInput(input string) ([]string, map[string]Path) {
	splitInput := strings.Split(input, "\n\n")
	directions := splitInput[0]
	pathStrings := strings.Split(splitInput[1], "\n")

	paths := make(map[string]Path)
	for _, pathString := range pathStrings {
		paths[pathString[0:3]] = Path{left: pathString[7:10], right: pathString[12:15]}
	}

	return strings.Split(directions, ""), paths
}

func finished(finishedAt []int) bool {
	for _, value := range finishedAt {
		if value == -1 {
			return false
		}
	}
	return true
}

func TestTestInput(t *testing.T) {
	directions, paths := parseInput(InputTest)
	var positions []string
	var finishedAt []int
	for key, _ := range paths {
		if key[2:3] == "A" {
			positions = append(positions, key)
			finishedAt = append(finishedAt, -1)
		}
	}

	steps := 0
	for !finished(finishedAt) {
		direction := directions[steps%len(directions)]
		if direction == "L" {
			for index, position := range positions {
				positions[index] = paths[position].left
			}
		} else {
			for index, position := range positions {
				positions[index] = paths[position].right
			}
		}
		steps++
		for index, position := range positions {
			if finishedAt[index] == -1 && position[2:3] == "Z" {
				finishedAt[index] = steps
			}
		}
	}

	fmt.Println(util.LCMArray(finishedAt))
}

func TestInput(t *testing.T) {
	directions, paths := parseInput(Input)
	var positions []string
	var finishedAt []int
	for key, _ := range paths {
		if key[2:3] == "A" {
			positions = append(positions, key)
			finishedAt = append(finishedAt, -1)
		}
	}

	steps := 0
	for !finished(finishedAt) {
		direction := directions[steps%len(directions)]
		if direction == "L" {
			for index, position := range positions {
				positions[index] = paths[position].left
			}
		} else {
			for index, position := range positions {
				positions[index] = paths[position].right
			}
		}
		steps++
		for index, position := range positions {
			if finishedAt[index] == -1 && position[2:3] == "Z" {
				finishedAt[index] = steps
			}
		}
	}

	fmt.Println(util.LCMArray(finishedAt))
}
