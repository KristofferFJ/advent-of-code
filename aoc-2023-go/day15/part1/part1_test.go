package part1

import (
	"fmt"
	"strings"
	"testing"
)

func TestInput(t *testing.T) {
	sum := 0
	for _, line := range strings.Split(Input, ",") {
		sum += hasValue(line)
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
