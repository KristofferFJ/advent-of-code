package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"strings"
	"testing"
)

func TestTestInput(t *testing.T) {
	var sequences [][]int
	input := strings.Split(InputTest, "\n")
	sum := 0
	for _, sequence := range input {
		sequences = append(sequences, util.IntArray(sequence))
	}
	for _, sequence := range sequences {
		sum += getPreviousValue(sequence)
	}

	if sum != 2 {
		t.Error()
	}
}

func TestInput(t *testing.T) {
	var sequences [][]int
	input := strings.Split(Input, "\n")
	sum := 0
	for _, sequence := range input {
		sequences = append(sequences, util.IntArray(sequence))
	}
	for _, sequence := range sequences {
		previousValue := getPreviousValue(sequence)
		sum += previousValue
		fmt.Printf("%d previous value %d\n", sequence, previousValue)
	}

	fmt.Println(sum)
}

func getPreviousValue(sequence []int) int {
	if allValuesEqual(sequence) {
		if len(sequence) == 0 {
			fmt.Println(sequence)
		}
		return sequence[0]
	}
	var subSequence []int
	for i := 0; i < len(sequence)-1; i++ {
		subSequence = append(subSequence, sequence[i+1]-sequence[i])
	}
	subSequencePreviousValue := getPreviousValue(subSequence)
	return sequence[0] - subSequencePreviousValue
}

func allValuesEqual(sequence []int) bool {
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] != sequence[i+1] {
			return false
		}
	}
	return true
}
