package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"strings"
	"testing"
)

type Springs struct {
	positions     []string
	configuration []int
}

func TestInput(t *testing.T) {
	var springList []Springs
	sum := 0
	for _, line := range strings.Split(Input, "\n") {
		positions := strings.Split(line, " ")[0]
		configuration := strings.Split(line, " ")[1]
		springList = append(springList,
			Springs{
				positions:     strings.Split(positions, ""),
				configuration: util.IntArray(configuration),
			})
	}

	for _, springs := range springList {
		sum += countValidSetups(springs)
	}

	fmt.Println(sum)
}

func countValidSetups(springs Springs) int {
	existingSprings := len(util.Filter(springs.positions, func(spring string) bool {
		return spring == "#"
	}))
	springsToPlace := util.SumIntArray(springs.configuration) - existingSprings
	positions := createPositions(springsToPlace, springs.positions, 0)
	validPositions := util.Filter(positions, func(positions []string) bool {
		return valid(Springs{configuration: springs.configuration, positions: positions})
	})

	return len(validPositions)
}

func createPositions(springsToPlace int, currentPosition []string, smallestIndex int) [][]string {
	if springsToPlace == 0 {
		return [][]string{currentPosition}
	}
	var result [][]string
	for i := smallestIndex; i < len(currentPosition); i++ {
		if currentPosition[i] == "?" {
			newPosition := util.Duplicate(currentPosition)
			newPosition[i] = "#"
			result = append(result, createPositions(springsToPlace-1, newPosition, i)...)
		}
	}
	return result
}

func valid(springs Springs) bool {
	var groups []int
	groupSize := 0
	for _, spring := range springs.positions {
		if spring == "#" {
			groupSize++
		} else if groupSize == 0 {
			continue
		} else {
			groups = append(groups, groupSize)
			groupSize = 0
		}
	}
	if groupSize != 0 {
		groups = append(groups, groupSize)
	}

	return util.IntArrayEqual(groups, springs.configuration)
}
