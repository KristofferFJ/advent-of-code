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

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	return (a / GCD(a, b)) * b
}

func LCMArray(numbers []int) int {
	lcm := 1
	for _, number := range numbers {
		lcm = LCM(number, lcm)
	}
	return lcm
}
