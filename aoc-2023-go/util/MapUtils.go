package util

func DuplicateMap[T any](m map[string]T) map[string]T {
	result := make(map[string]T)
	for key, value := range m {
		result[key] = value
	}
	return result
}

func AddIfAvailable[T any](key string, elems map[string][]T, elem T) {
	value, ok := elems[key]
	if ok {
		elems[key] = append(value, elem)
	}
}
