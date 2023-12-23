package util

import "reflect"

func Filter[T any](list []T, test func(T) bool) (filteredList []T) {
	for _, element := range list {
		if test(element) {
			filteredList = append(filteredList, element)
		}
	}
	return
}

func PointInList(point Point, list []Point) bool {
	for _, elem := range list {
		if elem.Row == point.Row && elem.Col == point.Col {
			return true
		}
	}
	return false
}

func Insert[T any](a []T, index int, value T) []T {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func RotateRight[T any](array [][]T) [][]T {
	var result [][]T
	rowLength := len(array[0])
	for i := 0; i < rowLength; i++ {
		var newRow []T
		for j := len(array) - 1; j >= 0; j-- {
			newRow = append(newRow, array[j][i])
		}
		result = append(result, newRow)
	}
	return result
}

func RotateLeft[T any](array [][]T) [][]T {
	var result [][]T
	rowLength := len(array[0])
	for i := rowLength - 1; i >= 0; i-- {
		var newRow []T
		for j := 0; j < len(array); j++ {
			newRow = append(newRow, array[j][i])
		}
		result = append(result, newRow)
	}
	return result
}

func IntArrayEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, elem := range a {
		if elem != b[i] {
			return false
		}
	}
	return true
}

func StringArrayEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, elem := range a {
		if elem != b[i] {
			return false
		}
	}
	return true
}

func StringArrayDifferences(a, b []string) int {
	if len(a) != len(b) {
		panic("Arrays must be of equal length")
	}
	sum := 0
	for i, elem := range a {
		if elem != b[i] {
			sum++
		}
	}
	return sum
}

func SumIntArray(a []int) int {
	sum := 0
	for _, elem := range a {
		sum += elem
	}
	return sum
}

func Duplicate[T any](list []T) []T {
	newList := make([]T, len(list))
	copy(newList, list)
	return newList
}

func RemoveDuplicates(slice [][]string) [][]string {
	unique := make([][]string, 0)
	for _, elem := range slice {
		if !contains(unique, elem) {
			unique = append(unique, elem)
		}
	}
	return unique
}

func contains(slice [][]string, item []string) bool {
	for _, v := range slice {
		if reflect.DeepEqual(v, item) {
			return true
		}
	}
	return false
}

func Contains[T comparable](list []T, elem T) bool {
	for _, element := range list {
		if element == elem {
			return true
		}
	}
	return false
}
