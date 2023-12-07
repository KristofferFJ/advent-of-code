package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"strconv"
	"testing"
)

func TestTestInput(t *testing.T) {
	sum := 0
	grid := util.ToGrid(InputTest)

	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		sum += evaluateLine(rowIndex, grid)
	}

	if sum != 4361 {
		t.Error()
	}
}

func TestInput(t *testing.T) {
	sum := 0
	grid := util.ToGrid(Input)

	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		sum += evaluateLine(rowIndex, grid)
	}

	fmt.Println(sum)
}

func evaluateLine(rowIndex int, grid [][]string) int {
	currentNumberString := ""
	startIndex := -1
	endIndex := -1
	sum := 0

	for index, symbol := range grid[rowIndex] {
		if util.IsNumber(symbol) {
			currentNumberString += symbol
			endIndex = index
			if startIndex == -1 {
				startIndex = index
			}
		} else if currentNumberString != "" {
			number, _ := strconv.Atoi(currentNumberString)
			if hasAdjacentSymbol(rowIndex, startIndex, endIndex, grid) {
				sum += number
			} else {
				fmt.Printf("In row %d, %d is NOT valid\n", rowIndex, number)
			}
			currentNumberString = ""
			startIndex = -1
		}
	}

	if currentNumberString != "" {
		lineEndingNumber, _ := strconv.Atoi(currentNumberString)
		if hasAdjacentSymbol(rowIndex, startIndex, endIndex, grid) {
			sum += lineEndingNumber
		}
	}

	fmt.Printf("Row %d add a sum of %d\n", rowIndex, sum)
	return sum
}

func hasAdjacentSymbol(
	rowIndex int,
	columnStartIndex int,
	columnEndIndex int,
	grid [][]string,
) bool {
	for row := rowIndex - 1; row <= rowIndex+1; row++ {
		for column := columnStartIndex - 1; column <= columnEndIndex+1; column++ {
			if row >= 0 && row < len(grid[0]) && column >= 0 && column < len(grid) {
				value := grid[row][column]
				if !util.IsNumber(value) && value != "." {
					return true
				}
			}
		}
	}

	return false
}
