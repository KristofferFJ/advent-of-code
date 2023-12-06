package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/internal"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"testing"
)

func TestTestInput(t *testing.T) {
	sum := 0
	lines := strings.Split(InputTest, "\n")
	for _, line := range lines {
		sum += evaluateCard(toCard(line))
	}

	if sum != 13 {
		t.Error()
	}
}

func TestInput(t *testing.T) {
	sum := 0
	lines := strings.Split(Input, "\n")
	for _, line := range lines {
		sum += evaluateCard(toCard(line))
	}

	fmt.Println(sum)
}

type Card struct {
	WinningNumbers []int
	Numbers        []int
}

func evaluateCard(card Card) int {
	value := 0
	for _, number := range card.Numbers {
		if slices.Contains(card.WinningNumbers, number) {
			if value == 0 {
				value = 1
			} else {
				value *= 2
			}
		}
	}
	return value
}

func toCard(string string) Card {
	cleaned := internal.Remove(string, regexp.MustCompile(`Card\s+\d+: `))
	split := strings.Split(cleaned, " | ")
	winningNumbers := internal.IntArray(split[0])
	numbers := internal.IntArray(split[1])
	return Card{
		WinningNumbers: winningNumbers,
		Numbers:        numbers,
	}
}

func TestParseRounds(t *testing.T) {
	testCases := []struct {
		input         string
		expectedCard  Card
		expectedValue int
	}{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			Card{
				[]int{41, 48, 83, 86, 17},
				[]int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			8,
		},
		{"Card 182: 11 18 63 73 64 39  9 92 82 62 |  8 27 69 64  3 53 73 11 21 39 10 18 35 44 56 62 75 72  4 51  6 42 82 37 76",
			Card{
				[]int{11, 18, 63, 73, 64, 39, 9, 92, 82, 62},
				[]int{8, 27, 69, 64, 3, 53, 73, 11, 21, 39, 10, 18, 35, 44, 56, 62, 75, 72, 4, 51, 6, 42, 82, 37, 76},
			},
			64,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actualCard := toCard(tc.input)
			if !reflect.DeepEqual(actualCard, tc.expectedCard) {
				t.Error()
			}
			if evaluateCard(actualCard) != tc.expectedValue {
				t.Error()
			}
		})
	}
}
