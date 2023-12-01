package part1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

func TestTestInput(t *testing.T) {
	lines := strings.Split(InputTest, "\n")
	result := 0
	for i := range lines {
		result += appendDigits(FindFirstInteger(lines[i]), FindLastInteger(lines[i]))
	}

	if result != 142 {
		t.Errorf("TestInput returns %d, expected %d", result, 142)
	}
}

func TestInput(t *testing.T) {
	lines := strings.Split(Input, "\n")
	result := 0
	for i := range lines {
		result += appendDigits(FindFirstInteger(lines[i]), FindLastInteger(lines[i]))
	}

	fmt.Println(result)
}

func appendDigits(digit1, digit2 int) int {
	appendedString := strconv.Itoa(digit1) + strconv.Itoa(digit2)
	result, _ := strconv.Atoi(appendedString)
	return result
}

func FindFirstInteger(input string) int {
	for _, char := range input {
		if unicode.IsDigit(char) {
			result, _ := strconv.Atoi(string(char))
			return result
		}
	}

	return 0
}

func FindLastInteger(input string) int {
	for _, char := range reverseString(input) {
		if unicode.IsDigit(char) {
			result, _ := strconv.Atoi(string(char))
			return result
		}
	}

	return 0
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func TestFindIntegers(t *testing.T) {
	testCases := []struct {
		input         string
		expectedFirst int
		expectedLast  int
	}{
		{"1abc2", 1, 2},
		{"pqr3stu8vwx", 3, 8},
		{"a1b2c3d4e5f", 1, 5},
		{"treb7uchet", 7, 7},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actualFirst := FindFirstInteger(tc.input)
			actualLast := FindLastInteger(tc.input)
			if actualFirst != tc.expectedFirst {
				t.Errorf("FindFirstInteger(%s): expected %d, got %d", tc.input, tc.expectedFirst, actualFirst)
			}
			if actualLast != tc.expectedLast {
				t.Errorf("FindLastInteger(%s): expected %d, got %d", tc.input, tc.expectedLast, actualLast)
			}
		})
	}
}
