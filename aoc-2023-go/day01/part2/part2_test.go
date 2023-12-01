package part1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestTestInput(t *testing.T) {
	lines := strings.Split(InputTest, "\n")
	result := 0
	for i := range lines {
		result += appendDigits(FindFirstNumberOrWord(lines[i]), FindLastNumberOrWord(lines[i]))
	}

	if result != 281 {
		t.Errorf("TestInput returns %d, expected %d", result, 281)
	}
}

func TestInput(t *testing.T) {
	lines := strings.Split(Input, "\n")
	result := 0
	for i := range lines {
		lineValue := appendDigits(FindFirstNumberOrWord(lines[i]), FindLastNumberOrWord(lines[i]))
		result += lineValue
	}

	fmt.Println(result)
}

func appendDigits(digit1, digit2 int) int {
	appendedString := strconv.Itoa(digit1) + strconv.Itoa(digit2)
	result, _ := strconv.Atoi(appendedString)
	return result
}

func FindFirstNumberOrWord(s string) int {
	pattern := `(\d|one|two|three|four|five|six|seven|eight|nine)`
	re := regexp.MustCompile(pattern)

	match := re.FindString(s)
	if match == "" {
		return 0
	}

	if num, err := strconv.Atoi(match); err == nil {
		return num
	}

	return wordToNumber(match)
}

func FindLastNumberOrWord(s string) int {
	highestIndex := -1
	value := 0
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i := range numbers {
		if strings.LastIndex(s, numbers[i]) > highestIndex {
			highestIndex = strings.LastIndex(s, numbers[i])
			value = wordToNumber(numbers[i])
		}
	}

	return value
}

func wordToNumber(word string) int {
	if num, err := strconv.Atoi(word); err == nil {
		return num
	}
	switch word {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	default:
		return 0
	}
}
func TestFindIntegers(t *testing.T) {
	testCases := []struct {
		input         string
		expectedFirst int
		expectedLast  int
	}{
		{"two1nine", 2, 9},
		{"eightwothree", 8, 3},
		{"abcone2threexyz", 1, 3},
		{"xtwone3four", 2, 4},
		{"4nineeightseven2", 4, 2},
		{"zoneight234", 1, 4},
		{"7pqrstsixteen", 7, 6},
		{"3fourtwofive6ksblffhpqoneightsz", 3, 8},
		{"seven2fournine4seven", 7, 7},
		{"3vdrzmnxp", 3, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actualFirst := FindFirstNumberOrWord(tc.input)
			actualLast := FindLastNumberOrWord(tc.input)
			if actualFirst != tc.expectedFirst {
				t.Errorf("FindFirstInteger(%s): expected %d, got %d", tc.input, tc.expectedFirst, actualFirst)
			}
			if actualLast != tc.expectedLast {
				t.Errorf("FindLastInteger(%s): expected %d, got %d", tc.input, tc.expectedLast, actualLast)
			}
		})
	}
}
