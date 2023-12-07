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

	if sum != 467835 {
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
			if util.InGrid(i, j, grid) && util.IsNumber(grid[i][j]) {
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
			if util.InGrid(i, j, grid) && util.IsNumber(grid[i][j]) {
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
		if util.IsNumber(grid[row][i]) {
			num = grid[row][i] + num
		} else {
			break
		}
	}
	for i := column + 1; i < len(grid[0]); i++ {
		if util.IsNumber(grid[row][i]) {
			num = num + grid[row][i]
		} else {
			break
		}
	}
	convertedNum, _ := strconv.Atoi(num)
	fmt.Printf("row=%d, column=%d, found=%d\n", row, column, convertedNum)
	return convertedNum
}
