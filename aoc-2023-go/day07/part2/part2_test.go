package part1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

const CARDS = "AKQT98765432J"

func sortHands(hands []string) {
	sort.SliceStable(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
	})
}

func compareHands(thisHand, thatHand string) bool {
	if assignValue(thisHand) > assignValue(thatHand) {
		return false
	}
	if assignValue(thisHand) < assignValue(thatHand) {
		return true
	}
	thisHandList := strings.Split(thisHand, "")
	thatHandList := strings.Split(thatHand, "")
	for i := 0; i < len(thisHandList); i++ {
		if strings.Index(CARDS, thisHandList[i]) < strings.Index(CARDS, thatHandList[i]) {
			return false
		}
		if strings.Index(CARDS, thisHandList[i]) > strings.Index(CARDS, thatHandList[i]) {
			return true
		}
	}
	return true
}

func assignValue(hand string) int {
	cards := strings.Split(strings.Split(hand, " ")[0], "")
	groupedCards := make(map[string]int)
	for _, card := range cards {
		groupedCards[card] += 1
	}

	jokers := groupedCards["J"]
	delete(groupedCards, "J")

	if jokers == 5 {
		return 7
	}

	for _, count := range groupedCards {
		if count+jokers == 5 {
			return 7
		}
	}
	for _, count := range groupedCards {
		if count+jokers == 4 {
			return 6
		}
	}

	numberOfThrees := 0
	numberOfTwos := 0
	numberOfOnes := 0

	for _, count := range groupedCards {
		if count == 3 {
			numberOfThrees++
		}
		if count == 2 {
			numberOfTwos++
		}
		if count == 1 {
			numberOfOnes++
		}
	}

	if (numberOfThrees == 1 && numberOfTwos == 1) || (numberOfTwos == 2 && jokers == 1) {
		return 5
	}

	for _, count := range groupedCards {
		if count+jokers == 3 {
			return 4
		}
	}

	if (numberOfTwos == 2) || (numberOfTwos == 1 && jokers == 1) {
		return 3
	}
	for _, count := range groupedCards {
		if count+jokers == 2 {
			return 2
		}
	}
	return 1
}

func TestTestInput(t *testing.T) {
	hands := strings.Split(InputTest, "\n")
	sortHands(hands)
	sum := 0
	for rank, hand := range hands {
		bet, _ := strconv.Atoi(strings.Split(hand, " ")[1])
		sum += (rank + 1) * bet
	}

	if sum != 5905 {
		t.Error()
	}
}

func TestInput(t *testing.T) {
	hands := strings.Split(Input, "\n")
	sortHands(hands)
	sum := 0
	for rank, hand := range hands {
		bet, _ := strconv.Atoi(strings.Split(hand, " ")[1])
		sum += (rank + 1) * bet
	}
	fmt.Println(sum)
}
