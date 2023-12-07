package util

import (
	"strconv"
)

func Max(one int, two int) int {
	if one > two {
		return one
	}
	return two
}

func Min(one int, two int) int {
	if one < two {
		return one
	}
	return two
}

func AppendNumbers(numbers []int) int {
	appended := ""
	for _, number := range numbers {
		appended += strconv.Itoa(number)
	}
	result, _ := strconv.Atoi(appended)
	return result
}
