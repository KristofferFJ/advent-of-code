package util

import (
	"regexp"
	"strconv"
	"strings"
)

func IsNumber(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

func Remove(string string, regexp *regexp.Regexp) string {
	return regexp.ReplaceAllString(string, "")
}

func IntArray(string string) []int {
	stringNumbers := regexp.MustCompile("-?\\d+").FindAllString(string, -1)
	var numbers []int
	for _, stringNumber := range stringNumbers {
		convertedNumber, _ := strconv.Atoi(stringNumber)
		numbers = append(numbers, convertedNumber)
	}
	return numbers
}

func Int(string string) int {
	number, _ := strconv.Atoi(string)
	return number
}

func Duplicate2D(rows, cols int, value string) string {
	var colList []string
	for _, line := range strings.Split(value, "\n") {
		duplicatedLine := line
		for i := 0; i < rows-1; i++ {
			duplicatedLine += line
		}
		colList = append(colList, duplicatedLine)
	}
	upperRow := strings.Join(colList, "\n")
	fullString := upperRow
	for i := 0; i < cols-1; i++ {
		fullString += "\n" + upperRow
	}

	return fullString
}
