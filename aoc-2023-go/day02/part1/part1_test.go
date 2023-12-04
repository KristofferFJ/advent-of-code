package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/internal"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type Rounds struct {
	Balls []Balls
}

type Balls struct {
	RedBalls   int
	BlueBalls  int
	GreenBalls int
}

func TestTestInput(t *testing.T) {
	actualBalls := Balls{RedBalls: 12, GreenBalls: 13, BlueBalls: 14}
	lines := strings.Split(InputTest, "\n")
	sum := 0
	for index, line := range lines {
		round := toRounds(line)
		if maxRed(round) <= actualBalls.RedBalls &&
			maxGreen(round) <= actualBalls.GreenBalls &&
			maxBlue(round) <= actualBalls.BlueBalls {
			sum += index + 1
		}
	}

	if sum != 8 {
		t.Error()
	}
}

func TestInput(t *testing.T) {
	actualBalls := Balls{RedBalls: 12, GreenBalls: 13, BlueBalls: 14}
	lines := strings.Split(Input, "\n")
	sum := 0
	for index, line := range lines {
		round := toRounds(line)
		if maxRed(round) <= actualBalls.RedBalls &&
			maxGreen(round) <= actualBalls.GreenBalls &&
			maxBlue(round) <= actualBalls.BlueBalls {
			sum += index + 1
		}
	}
	fmt.Println(sum)
}

func maxRed(round Rounds) int {
	maxRed := 0
	for _, balls := range round.Balls {
		if balls.RedBalls > maxRed {
			maxRed = balls.RedBalls
		}
	}
	return maxRed
}

func maxBlue(round Rounds) int {
	maxBlue := 0
	for _, balls := range round.Balls {
		if balls.BlueBalls > maxBlue {
			maxBlue = balls.BlueBalls
		}
	}
	return maxBlue
}

func maxGreen(round Rounds) int {
	maxGreen := 0
	for _, balls := range round.Balls {
		if balls.GreenBalls > maxGreen {
			maxGreen = balls.GreenBalls
		}
	}
	return maxGreen
}

func toRounds(string string) Rounds {
	cleaned := internal.Remove(string, regexp.MustCompile(`Game \d+: `))
	roundsStrings := strings.Split(cleaned, "; ")

	rounds := Rounds{Balls: []Balls{}}
	for _, round := range roundsStrings {
		rounds.Balls = append(rounds.Balls, toBalls(round))
	}

	return rounds
}

func toBalls(roundString string) Balls {
	balls := Balls{RedBalls: 0, BlueBalls: 0, GreenBalls: 0}
	ballSplit := strings.Split(roundString, ", ")
	for _, ball := range ballSplit {
		numberColorSplit := strings.Split(ball, " ")
		number, _ := strconv.Atoi(numberColorSplit[0])
		color := numberColorSplit[1]
		switch color {
		case "green":
			balls.GreenBalls = number
		case "red":
			balls.RedBalls = number
		case "blue":
			balls.BlueBalls = number
		}
	}

	return balls
}

func TestParseRounds(t *testing.T) {
	testCases := []struct {
		input          string
		expectedRounds Rounds
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", Rounds{
			[]Balls{
				{BlueBalls: 3, GreenBalls: 0, RedBalls: 4},
				{BlueBalls: 6, GreenBalls: 2, RedBalls: 1},
				{BlueBalls: 0, GreenBalls: 2, RedBalls: 0},
			},
		}},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", Rounds{
			[]Balls{
				{BlueBalls: 1, GreenBalls: 2, RedBalls: 0},
				{BlueBalls: 4, GreenBalls: 3, RedBalls: 1},
				{BlueBalls: 1, GreenBalls: 1, RedBalls: 0},
			},
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actualRounds := toRounds(tc.input)
			if !reflect.DeepEqual(actualRounds, tc.expectedRounds) {
				t.Error()
			}
		})
	}
}
