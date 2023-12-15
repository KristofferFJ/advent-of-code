package part1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestInput(t *testing.T) {
	boxes := make(map[int][]string)
	operationRegex := regexp.MustCompile("[=\\-]")
	for _, line := range strings.Split(Input, ",") {
		operationIndex := operationRegex.FindStringIndex(line)[0]
		operation := string(line[operationIndex])
		box := hasValue(line[:operationIndex])
		if operation == "=" {
			lenses := boxes[box]
			replaced := false
			newLens := strings.ReplaceAll(line, "=", " ")
			for i := 0; i < len(lenses); i++ {
				if strings.Contains(lenses[i], line[:operationIndex]) {
					boxes[box][i] = newLens
					replaced = true
					break
				}
			}
			if !replaced {
				boxes[box] = append(boxes[box], newLens)
			}
		}
		if operation == "-" {
			lenses := boxes[box]
			for i := 0; i < len(lenses); i++ {
				if strings.Contains(lenses[i], line[:operationIndex]) {
					boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
					break
				}
			}
		}
	}
	sumBox(boxes)
}

func sumBox(box map[int][]string) {
	sum := 0
	for key, value := range box {
		for index, line := range value {
			focalLength, _ := strconv.Atoi(strings.Split(line, " ")[1])
			sum += (key + 1) * (index + 1) * focalLength
		}
	}
	fmt.Println(sum)
}

func hasValue(hash string) int {
	value := 0
	for _, char := range hash {
		value += int(char)
		value *= 17
		value %= 256
	}
	return value
}
