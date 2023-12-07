package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"testing"
)

func TestTestInput(t *testing.T) {
	sum := 0
	lines := strings.Split(InputTest, "\n")
	var cardCollections []*CardCollection

	for _, line := range lines {
		cardCollections = append(cardCollections, &CardCollection{Amount: 1, Card: toCard(line)})
	}

	for index, cardCollection := range cardCollections {
		score := evaluateCard(cardCollection.Card)
		for i := index + 1; i <= index+score; i++ {
			if len(cardCollections) > i {
				cardCollections[i].Amount += cardCollection.Amount
			}
		}
	}

	for _, cardCollection := range cardCollections {
		sum += cardCollection.Amount
	}

	if sum != 30 {
		t.Error()
	}
}

func TestInput(t *testing.T) {
	sum := 0
	lines := strings.Split(Input, "\n")
	var cardCollections []*CardCollection

	for _, line := range lines {
		cardCollections = append(cardCollections, &CardCollection{Amount: 1, Card: toCard(line)})
	}

	for index, cardCollection := range cardCollections {
		score := evaluateCard(cardCollection.Card)
		for i := index + 1; i <= index+score; i++ {
			if len(cardCollections) > i {
				cardCollections[i].Amount += cardCollection.Amount
			}
		}
	}

	for _, cardCollection := range cardCollections {
		sum += cardCollection.Amount
	}

	fmt.Println(sum)
}

type CardCollection struct {
	Card   Card
	Amount int
}

type Card struct {
	WinningNumbers []int
	Numbers        []int
}

func evaluateCard(card Card) int {
	value := 0
	for _, number := range card.Numbers {
		if slices.Contains(card.WinningNumbers, number) {
			value += 1
		}
	}
	return value
}

func toCard(string string) Card {
	cleaned := Remove(string, regexp.MustCompile(`Card\s+\d+: `))
	split := strings.Split(cleaned, " | ")
	winningNumbers := util.IntArray(split[0])
	numbers := util.IntArray(split[1])
	return Card{
		WinningNumbers: winningNumbers,
		Numbers:        numbers,
	}
}

func Remove(string string, regexp *regexp.Regexp) string {
	return regexp.ReplaceAllString(string, "")
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
			4,
		},
		{"Card 182: 11 18 63 73 64 39  9 92 82 62 |  8 27 69 64  3 53 73 11 21 39 10 18 35 44 56 62 75 72  4 51  6 42 82 37 76",
			Card{
				[]int{11, 18, 63, 73, 64, 39, 9, 92, 82, 62},
				[]int{8, 27, 69, 64, 3, 53, 73, 11, 21, 39, 10, 18, 35, 44, 56, 62, 75, 72, 4, 51, 6, 42, 82, 37, 76},
			},
			7,
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
