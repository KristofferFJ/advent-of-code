package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/internal"
	"strings"
	"testing"
)

func TestTestInput(t *testing.T) {
	seedIntervals := seedIntervals(InputTest)
	fmt.Println(seedIntervals)
}

func TestInput(t *testing.T) {

}

func getLocationFromSeed(seedRanges []interval, maps destinationSourceMaps) int {
	soil := getCorresponding(seedRanges, maps.seedToSoil)
	fertilizer := getCorresponding(soil, maps.soilToFertilizer)
	water := getCorresponding(fertilizer, maps.fertilizerToWater)
	light := getCorresponding(water, maps.waterToLight)
	temperature := getCorresponding(light, maps.lightToTemperature)
	humidity := getCorresponding(temperature, maps.temperatureToHumidity)
	return getCorresponding(humidity, maps.humidityToLocation)
}

func getCorresponding(intervals []interval, destinationSourceMap destinationSourceMap) []interval {

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
	sourceInterval []interval
	modifier       int
}

func seedIntervals(input string) []interval {
	numbers := internal.IntArray(strings.Split(input, "\n\n")[0])
	var intervals []interval
	for startIndex := 0; startIndex < len(numbers); startIndex += 2 {
		intervals = append(intervals, interval{
			startInclusive: numbers[startIndex],
			endInclusive:   numbers[startIndex] + numbers[startIndex+1] - 1,
		})
	}

	return intervals
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
		numbers := internal.IntArray(split[i])
		destinationSourceMaps = append(destinationSourceMaps, destinationSourceMap{
			destination: numbers[0],
			source:      numbers[1],
			length:      numbers[2],
		})
	}
	return destinationSourceMaps
}
