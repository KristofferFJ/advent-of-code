package internal

import (
	"regexp"
	"strconv"
)

func IsNumber(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

func Remove(string string, regexp *regexp.Regexp) string {
	return regexp.ReplaceAllString(string, "")
}
