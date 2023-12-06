package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/internal"
	"strings"
	"testing"
)

func TestTestInput(t *testing.T) {
	sum := 0
	raceTime := internal.AppendNumbers(internal.FindNumbersInString(strings.Split(InputTest, "\n")[0]))
	distanceToBeat := internal.AppendNumbers(internal.FindNumbersInString(strings.Split(InputTest, "\n")[1]))

	for holdTime := 1; holdTime < raceTime; holdTime++ {
		if getDistanceTravelled(holdTime, raceTime) > distanceToBeat {
			sum += 1
		}
	}

	if sum != 71503 {
		t.Error()
	}
}

func TestInput(t *testing.T) {
	sum := 0
	raceTime := internal.AppendNumbers(internal.FindNumbersInString(strings.Split(Input, "\n")[0]))
	distanceToBeat := internal.AppendNumbers(internal.FindNumbersInString(strings.Split(Input, "\n")[1]))

	for holdTime := 1; holdTime < raceTime; holdTime++ {
		if getDistanceTravelled(holdTime, raceTime) > distanceToBeat {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func getDistanceTravelled(holdTime int, raceTime int) int {
	return (raceTime - holdTime) * holdTime
}
