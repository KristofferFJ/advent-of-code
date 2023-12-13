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

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Binomial(n, k int) int {
	if k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	return Binomial(n-1, k-1) + Binomial(n-1, k)
}

func ToBinary(num int, length int) []int {
	var result []int
	for num > 0 {
		result = append(result, num%2)
		num /= 2
	}

	prepended := make([]int, length)
	copy(prepended[length-len(result):], result)
	return prepended
}
