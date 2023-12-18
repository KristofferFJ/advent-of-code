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
			})
	}

	for i := 0; i < len(setups); i++ {
		sum += countValidSetups(setups[i])
	}

	fmt.Println(sum)
}

func key(springs string, config []int) string {
	return springs + fmt.Sprintf("%v", config)
}

var cache = make(map[string]int)

func countValidSetups(setup Setup) int {
	cache = make(map[string]int)
	value := countValidSetupsRecursive(setup.springs, setup.configuration, 0)
	fmt.Printf("%s: %d\n", setup.springs, value)
	return value
}

func countValidSetupsRecursive(spring string, config []int, index int) int {
	if strings.Index(spring, "?") == -1 {
		if valid(spring, config) {
			return 1
		}
		return 0
	}

	count, ok := cache[key(spring, config)]
	if ok {
		return count
	} else {
		validConfig, remainingSpring, remainingConfig := possiblyValid(spring, config)
		index -= len(spring) - len(remainingSpring)
		if index > 0 {
			spring = remainingSpring
			config = remainingConfig
		} else {
			index += len(spring) - len(remainingSpring)
		}
		if !validConfig {
			cache[key(spring, config)] = 0
			return 0
		}
	}

	next := spring[index : index+1]
	validSetups := 0
	if next == "#" {
		validSetups = countValidSetupsRecursive(spring, config, index+1)
	}
	if next == "." {
		validSetups = countValidSetupsRecursive(spring, config, index+1)
	}
	if next == "?" {
		validSetups = countValidSetupsRecursive(spring[:index]+"#"+spring[index+1:], config, index+1) +
			countValidSetupsRecursive(spring[:index]+"."+spring[index+1:], config, index+1)
	}

	cache[key(spring, config)] = validSetups
	return validSetups
}

func possiblyValid(springs string, configuration []int) (bool, string, []int) {
	finishedUpUntil := strings.Index(springs, "?")
	remainingSpring := springs
	remainingConfig := configuration

	springsArray := strings.Split(springs, "")
	if finishedUpUntil > 0 {
		var finishedGroups []int

		for _, segment := range strings.Split(springs[:(finishedUpUntil)], ".") {
			if segment != "" {
				finishedGroups = append(finishedGroups, len(segment))
			}
		}

		if len(finishedGroups) > len(configuration) {
			return false, springs, configuration
		}

		for i := 0; i < len(finishedGroups)-1; i++ {
			if finishedGroups[i] != configuration[i] {
				return false, springs, configuration
			}
		}

		if len(finishedGroups) > 0 {
			if finishedGroups[len(finishedGroups)-1] > configuration[len(finishedGroups)-1] {
				return false, springs, configuration
			}
		}

		minFirstGroup := 0
		start := util.Min(strings.Index(springs, "#"), strings.Index(springs, "?"))
		if start != -1 {
			for i := start; i < len(springsArray); i++ {
				if springsArray[i] == "#" {
					minFirstGroup += 1
				} else {
					break
				}
			}
		}

		if len(configuration) > 0 && minFirstGroup > configuration[0] {
			return false, springs, configuration
		}

		if finishedGroups != nil && springs[finishedUpUntil-1:finishedUpUntil] == "#" {
			finishedGroups = finishedGroups[:len(finishedGroups)-1]
		}

		if finishedGroups != nil {
			for finishedUpUntil > 0 {
				if springs[finishedUpUntil-1:finishedUpUntil] == "." {
					break
				}
				finishedUpUntil -= 1
			}
			remainingSpring = springs[finishedUpUntil:]
			remainingConfig = configuration[len(finishedGroups):]
		}
	}

	return true, remainingSpring, remainingConfig
}

func valid(springs string, configuration []int) bool {
	segments := strings.Split(springs, ".")
	var counts []int

	for _, segment := range segments {
		if segment == "" {
			continue
		}
		counts = append(counts, len(segment))
	}

	return util.IntArrayEqual(counts, configuration)
}
