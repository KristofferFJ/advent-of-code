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
		sum += getNextValue(sequence)
	}

	if sum != 114 {
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
		nextValue := getNextValue(sequence)
		sum += nextValue
		fmt.Printf("%d next value %d\n", sequence, nextValue)
	}

	fmt.Println(sum)
}

func getNextValue(sequence []int) int {
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
	subSequenceNextValue := getNextValue(subSequence)
	return sequence[len(sequence)-1] + subSequenceNextValue
}

func allValuesEqual(sequence []int) bool {
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] != sequence[i+1] {
			return false
		}
	}
	return true
}
