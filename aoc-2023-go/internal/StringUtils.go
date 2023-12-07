package internal

import (
	"regexp"
	"strconv"
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
	stringNumbers := regexp.MustCompile("\\d+").FindAllString(string, -1)
	var numbers []int
	for _, stringNumber := range stringNumbers {
		convertedNumber, _ := strconv.Atoi(stringNumber)
		numbers = append(numbers, convertedNumber)
	}
	return numbers
}
