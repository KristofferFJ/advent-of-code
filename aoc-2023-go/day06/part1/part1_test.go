package part1

import (
	"fmt"
	. "io.kristofferfj.github/aoc-2023-go/internal"
	"strings"
	"testing"
)

func TestTestInput(t *testing.T) {
	product := 1
	raceTimes := IntArray(strings.Split(InputTest, "\n")[0])
	distancesToBeat := IntArray(strings.Split(InputTest, "\n")[1])

	for index, raceTime := range raceTimes {
		sum := 0
		for holdTime := 1; holdTime < raceTime; holdTime++ {
			if getDistanceTravelled(holdTime, raceTime) > distancesToBeat[index] {
				sum += 1
			}
		}
		product *= sum
	}

	if product != 288 {
		t.Error()
	}
}

func TestInput(t *testing.T) {
	product := 1
	raceTimes := IntArray(strings.Split(Input, "\n")[0])
	distancesToBeat := IntArray(strings.Split(Input, "\n")[1])

	for index, raceTime := range raceTimes {
		sum := 0
		for holdTime := 1; holdTime < raceTime; holdTime++ {
			if getDistanceTravelled(holdTime, raceTime) > distancesToBeat[index] {
				sum += 1
			}
		}
		product *= sum
	}

	fmt.Println(product)
}

func getDistanceTravelled(holdTime int, raceTime int) int {
	return (raceTime - holdTime) * holdTime
}
