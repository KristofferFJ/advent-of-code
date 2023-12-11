package util

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
		if elem.Row == point.Row && elem.Column == point.Column {
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
