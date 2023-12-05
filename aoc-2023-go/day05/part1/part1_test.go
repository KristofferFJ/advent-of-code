package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/internal"
	"strings"
	"testing"
)

func TestTestInput(t *testing.T) {
	seedNumbers := seedNumbers(InputTest)
	parsed := parseInput(InputTest)
	for _, seed := range seedNumbers {
		fmt.Println(getLocationFromSeed(seed, parsed))
	}
}

func TestInput(t *testing.T) {
	seedNumbers := seedNumbers(Input)
	parsed := parseInput(Input)
	lowest := -1
	for _, seed := range seedNumbers {
		location := getLocationFromSeed(seed, parsed)
		if lowest == -1 || location < lowest {
			lowest = location
		}
	}
	println(lowest)
}

func getLocationFromSeed(seed int, maps destinationSourceMaps) int {
	soil := getCorresponding(seed, maps.seedToSoil)
	fertilizer := getCorresponding(soil, maps.soilToFertilizer)
	water := getCorresponding(fertilizer, maps.fertilizerToWater)
	light := getCorresponding(water, maps.waterToLight)
	temperature := getCorresponding(light, maps.lightToTemperature)
	humidity := getCorresponding(temperature, maps.temperatureToHumidity)
	return getCorresponding(humidity, maps.humidityToLocation)
}

func getCorresponding(start int, destinationSourceMap []destinationSourceMap) int {
	for _, seedToSoil := range destinationSourceMap {
		if start >= seedToSoil.source && start < seedToSoil.source+seedToSoil.length {
			return seedToSoil.destination + start - seedToSoil.source
		}
	}
	return start
}

type destinationSourceMaps struct {
	seedToSoil            []destinationSourceMap
	soilToFertilizer      []destinationSourceMap
	fertilizerToWater     []destinationSourceMap
	waterToLight          []destinationSourceMap
	lightToTemperature    []destinationSourceMap
	temperatureToHumidity []destinationSourceMap
	humidityToLocation    []destinationSourceMap
}

type destinationSourceMap struct {
	destination int
	source      int
	length      int
}

func seedNumbers(input string) []int {
	return internal.FindNumbersInString(strings.Split(input, "\n\n")[0])
}

func parseInput(input string) destinationSourceMaps {
	partitions := strings.Split(input, "\n\n")
	return destinationSourceMaps{
		seedToSoil:            toDestinationSourceMaps(partitions[1]),
		soilToFertilizer:      toDestinationSourceMaps(partitions[2]),
		fertilizerToWater:     toDestinationSourceMaps(partitions[3]),
		waterToLight:          toDestinationSourceMaps(partitions[4]),
		lightToTemperature:    toDestinationSourceMaps(partitions[5]),
		temperatureToHumidity: toDestinationSourceMaps(partitions[6]),
		humidityToLocation:    toDestinationSourceMaps(partitions[7]),
	}
}

func toDestinationSourceMaps(partition string) []destinationSourceMap {
	split := strings.Split(partition, "\n")
	var destinationSourceMaps []destinationSourceMap
	for i := 1; i < len(split); i++ {
		numbers := internal.FindNumbersInString(split[i])
		destinationSourceMaps = append(destinationSourceMaps, destinationSourceMap{
			destination: numbers[0],
			source:      numbers[1],
			length:      numbers[2],
		})
	}
	return destinationSourceMaps
}
