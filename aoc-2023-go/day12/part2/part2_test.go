package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"strings"
	"testing"
)

type Setup struct {
	springs       string
	configuration []int
	config        string
}

func TestInput(t *testing.T) {
	var setups []Setup
	sum := 0
	for _, line := range strings.Split(InputTest, "\n") {
		springs := strings.Split(line, " ")[0]
		alteredSprings := springs
		for i := 0; i < 4; i++ {
			alteredSprings += "?" + springs
		}
		config := strings.Split(line, " ")[1]
		alteredConfig := config
		for i := 0; i < 4; i++ {
			alteredConfig += "," + config
		}

		setups = append(setups,
			Setup{
				springs:       alteredSprings,
				configuration: util.IntArray(alteredConfig),
				config:        alteredConfig,
			})
	}

	for _, springs := range setups {
		sum += countValidSetups(springs)
	}

	fmt.Println(sum)
}
func countValidSetups(setup Setup) int {
	springCount := strings.Count(setup.springs, "#")
	questionMarkCount := strings.Count(setup.springs, "?")
	springsToPlace := util.SumIntArray(setup.configuration) - springCount
	dotsToPlace := questionMarkCount - springsToPlace

	return len(getValidPositions(springsToPlace, dotsToPlace, setup.springs, setup.config))
}

func getValidPositions(springsToPlace, dotsToPlace int, springs, config string) []string {
	if springsToPlace == 0 && dotsToPlace == 0 {
		if valid(springs, config) {
			return []string{springs}
		}
		return []string{}
	}

	if largestSegment(springs) > util.MaxArray(util.IntArray(config)) {
		return []string{}
	}

	if springsToPlace > 0 && dotsToPlace > 0 {
		return append(getValidPositions(springsToPlace-1, dotsToPlace, strings.Replace(springs, "?", "#", 1), config),
			getValidPositions(springsToPlace, dotsToPlace-1, strings.Replace(springs, "?", ".", 1), config)...)
	}
	if springsToPlace > 0 {
		return getValidPositions(springsToPlace-1, dotsToPlace, strings.Replace(springs, "?", "#", 1), config)
	}
	return getValidPositions(springsToPlace, dotsToPlace-1, strings.Replace(springs, "?", ".", 1), config)
}

func valid(springs string, config string) bool {
	segments := strings.Split(springs, ".")
	var counts []string

	for _, segment := range segments {
		if segment == "" {
			continue
		}
		counts = append(counts, fmt.Sprintf("%d", len(segment)))
	}

	return strings.Join(counts, ",") == config
}

func largestSegment(springs string) int {
	segments := strings.Split(springs, ".")
	largest := 0
	for _, segment := range segments {
		if len(segment) > largest {
			largest = len(segment)
		}
	}
	return largest

}
