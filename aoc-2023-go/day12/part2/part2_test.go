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
	for _, line := range strings.Split(InputTest, "\n") {
		positions := strings.Split(strings.Split(line, " ")[0], "")
		alteredPositions := util.Duplicate(positions)
		for i := 0; i < 4; i++ {
			alteredPositions = append(alteredPositions, "?")
			alteredPositions = append(alteredPositions, positions...)
		}
		configuration := util.IntArray(strings.Split(line, " ")[1])
		alteredConfiguration := util.Duplicate(configuration)
		for i := 0; i < 4; i++ {
			alteredConfiguration = append(alteredConfiguration, configuration...)
		}
		springList = append(springList,
			Springs{
				positions:     alteredPositions,
				configuration: alteredConfiguration,
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
	positions := createPositions(springsToPlace, springs.positions, 0, springs.configuration)
	validPositions := util.Filter(positions, func(positions []string) bool {
		return valid(Springs{configuration: springs.configuration, positions: positions})
	})

	return len(validPositions)
}

func createPositions(springsToPlace int, currentPosition []string, smallestIndex int, configuration []int) [][]string {
	if springsToPlace == 0 {
		return [][]string{currentPosition}
	}
	var result [][]string
	for i := smallestIndex; i < len(currentPosition); i++ {
		if currentPosition[i] == "?" {
			newPosition := util.Duplicate(currentPosition)
			newPosition[i] = "#"
			if !potentiallyValid(newPosition, configuration) {
				continue
			}
			result = append(result, createPositions(springsToPlace-1, newPosition, i, configuration)...)
		}
	}
	return result
}

func potentiallyValid(position []string, configuration []int) bool {
	var groups []int
	groupSize := 0
	for _, spring := range position {
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

	for i := 0; i < util.Min(len(groups), len(configuration)); i++ {
		if groups[i] > configuration[i] {
			return false
		}
	}
	return true
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
