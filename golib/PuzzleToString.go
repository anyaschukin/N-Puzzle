package golib

import (
	"strconv"
	"strings"
)

// PuzzleToString converts a Puzzle []int to a string
func PuzzleToString(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}
