package part1

import (
	"fmt"
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

func TestTestInput(t *testing.T) {
	directions, paths := parseInput(InputTest)
	position := "AAA"
	steps := 0
	for position != "ZZZ" {
		direction := directions[steps%len(directions)]
		if direction == "L" {
			position = paths[position].left
		} else {
			position = paths[position].right
		}
		steps++
	}

	fmt.Println(steps)
}

func TestInput(t *testing.T) {
	directions, paths := parseInput(Input)
	position := "AAA"
	steps := 0
	for position != "ZZZ" {
		direction := directions[steps%len(directions)]
		if direction == "L" {
			position = paths[position].left
		} else {
			position = paths[position].right
		}
		steps++
	}

	fmt.Println(steps)
}
