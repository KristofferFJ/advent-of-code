package part1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestTestInput(t *testing.T) {
	sum := 0
	grid := toGrid(InputTest)

	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		sum += evaluateLine(rowIndex, grid)
	}

	if sum != 467835 {
		t.Error()
	}
}

func TestInput(t *testing.T) {
	sum := 0
	grid := toGrid(Input)

	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		sum += evaluateLine(rowIndex, grid)
	}

	fmt.Println(sum)
}

func evaluateLine(rowIndex int, grid [][]string) int {
	sum := 0

	for columnIndex, symbol := range grid[rowIndex] {
		if symbol == "*" {
			if numberOfAdjacentNumbers(rowIndex, columnIndex, grid) == 2 {
				sum += calculateGearRatio(rowIndex, columnIndex, grid)

			}
		}
	}

	return sum
}

func numberOfAdjacentNumbers(
	row int,
	column int,
	grid [][]string,
) int {
	sum := 0
	number := false

	for i := row - 1; i <= row+1; i++ {
		number = false
		for j := column - 1; j <= column+1; j++ {
			if inGrid(i, j, grid) && IsNumber(grid[i][j]) {
				if !number {
					sum += 1
					number = true
				} else {
					number = true
				}
			} else {
				number = false
			}
		}
	}

	return sum
}

func calculateGearRatio(
	row int,
	column int,
	grid [][]string,
) int {
	product := 1
	number := false

	for i := row - 1; i <= row+1; i++ {
		number = false
		for j := column - 1; j <= column+1; j++ {
			if inGrid(i, j, grid) && IsNumber(grid[i][j]) {
				if !number {
					product *= getNumber(i, j, grid)
					number = true
				} else {
					number = true
				}
			} else {
				number = false
			}
		}
	}

	fmt.Printf("product=%d\n", product)
	return product
}

func getNumber(row int, column int, grid [][]string) int {
	num := grid[row][column]
	for i := column - 1; i >= 0; i-- {
		if IsNumber(grid[row][i]) {
			num = grid[row][i] + num
		} else {
			break
		}
	}
	for i := column + 1; i < len(grid[0]); i++ {
		if IsNumber(grid[row][i]) {
			num = num + grid[row][i]
		} else {
			break
		}
	}
	convertedNum, _ := strconv.Atoi(num)
	fmt.Printf("row=%d, column=%d, found=%d\n", row, column, convertedNum)
	return convertedNum
}

func inGrid(row int, column int, grid [][]string) bool {
	return row >= 0 && row < len(grid[0]) && column >= 0 && column < len(grid)
}

func IsNumber(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

func toGrid(s string) [][]string {
	lines := strings.Split(s, "\n")
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}
