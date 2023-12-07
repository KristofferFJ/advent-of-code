package part1

import (
	"fmt"
	. "io.kristofferfj.github/aoc-2023-go/util"
	"strings"
	"testing"
)

func TestTestInput(t *testing.T) {
	seedIntervals := seedIntervals(InputTest)
	destinationSourceMaps := parseInput(InputTest)

	finalIntervals := getLocationFromSeed(seedIntervals, destinationSourceMaps)

	if finalIntervals[0].StartInclusive != 46 {
		t.Error()
	}
}

func TestInput(t *testing.T) {
	seedIntervals := seedIntervals(Input)
	destinationSourceMaps := parseInput(Input)

	finalIntervals := getLocationFromSeed(seedIntervals, destinationSourceMaps)

	fmt.Println(finalIntervals[0].StartInclusive)
}

func getLocationFromSeed(seedRanges []Interval, maps destinationSourceMaps) []Interval {
	soil := getCorresponding(seedRanges, maps.seedToSoil)
	fertilizer := getCorresponding(soil, maps.soilToFertilizer)
	water := getCorresponding(fertilizer, maps.fertilizerToWater)
	light := getCorresponding(water, maps.waterToLight)
	temperature := getCorresponding(light, maps.lightToTemperature)
	humidity := getCorresponding(temperature, maps.temperatureToHumidity)
	return getCorresponding(humidity, maps.humidityToLocation)
}

func getCorresponding(intervals []Interval, destinationSourceMap []destinationSourceMap) []Interval {
	var modifiedIntervals []Interval
	var intersections []Interval
	for _, interval := range intervals {
		for _, sourceInterval := range destinationSourceMap {
			intersection, intersects := interval.Intersection(sourceInterval.sourceInterval)
			if intersects {
				intersections = append(intersections, intersection)
				modifiedIntervals = append(modifiedIntervals, Interval{
					StartInclusive: intersection.StartInclusive + sourceInterval.modifier,
					EndInclusive:   intersection.EndInclusive + sourceInterval.modifier,
				})
			}
		}
	}
	modifiedIntervals = append(modifiedIntervals, RemoveSlices(intervals, intersections)...)
	modifiedIntervals = MergeIntervals(modifiedIntervals)
	return modifiedIntervals
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
	sourceInterval Interval
	modifier       int
}

func seedIntervals(input string) []Interval {
	numbers := IntArray(strings.Split(input, "\n\n")[0])
	var intervals []Interval
	for startIndex := 0; startIndex < len(numbers); startIndex += 2 {
		intervals = append(intervals, Interval{
			StartInclusive: numbers[startIndex],
			EndInclusive:   numbers[startIndex] + numbers[startIndex+1] - 1,
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
		numbers := IntArray(split[i])
		destinationSourceMaps = append(destinationSourceMaps, destinationSourceMap{
			sourceInterval: Interval{
				StartInclusive: numbers[1],
				EndInclusive:   numbers[1] + numbers[2] - 1,
			},
			modifier: numbers[0] - numbers[1],
		})
	}
	return destinationSourceMaps
}
