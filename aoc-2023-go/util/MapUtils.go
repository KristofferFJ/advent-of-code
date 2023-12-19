package util

func DuplicateMap[S string, T any](m map[string]T) map[string]T {
	result := make(map[string]T)
	for key, value := range m {
		result[key] = value
	}
	return result
}
