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
	for _, line := range strings.Split(Input, "\n") {
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

	for i := 0; i < len(setups); i++ {
		sum += countValidSetups(setups[i])
	}

	fmt.Println(sum)
}
func countValidSetups(setup Setup) int {
	springCount := strings.Count(setup.springs, "#")
	questionMarkCount := strings.Count(setup.springs, "?")
	springsToPlace := util.SumIntArray(setup.configuration) - springCount
	dotsToPlace := questionMarkCount - springsToPlace

	return len(getValidPositions(springsToPlace, dotsToPlace, setup.springs, setup.config, setup.configuration))
}

func getValidPositions(springsToPlace, dotsToPlace int, springs, config string, configuration []int) []string {
	if springsToPlace == 0 && dotsToPlace == 0 {
		if valid(springs, config) {
			return []string{springs}
		}
		return []string{}
	}

	if !possiblyValid(springs, configuration) {
		return []string{}
	}

	if springsToPlace > 0 && dotsToPlace > 0 {
		return append(getValidPositions(springsToPlace-1, dotsToPlace, strings.Replace(springs, "?", "#", 1), config, configuration),
			getValidPositions(springsToPlace, dotsToPlace-1, strings.Replace(springs, "?", ".", 1), config, configuration)...)
	}
	if springsToPlace > 0 {
		return getValidPositions(springsToPlace-1, dotsToPlace, strings.Replace(springs, "?", "#", 1), config, configuration)
	}
	return getValidPositions(springsToPlace, dotsToPlace-1, strings.Replace(springs, "?", ".", 1), config, configuration)
}

func possiblyValid(springs string, configuration []int) bool {
	finishedUpUntil := strings.Index(springs, "?")
	if finishedUpUntil > 0 {
		var finishedGroups []int
		for _, segment := range strings.Split(springs[:(finishedUpUntil)], ".") {
			if segment != "" {
				finishedGroups = append(finishedGroups, len(segment))
			}
		}

		for i := 0; i < len(finishedGroups)-1; i++ {
			if finishedGroups[i] != configuration[i] {
				return false
			}
		}
		if len(finishedGroups) > 0 && finishedGroups[len(finishedGroups)-1] > configuration[len(finishedGroups)-1] {
			return false
		}
		if len(finishedGroups) > 0 && springs[finishedUpUntil-1:finishedUpUntil] == "." && finishedGroups[len(finishedGroups)-1] != configuration[len(finishedGroups)-1] {
			return false
		}
	}

	return true
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
