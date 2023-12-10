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
