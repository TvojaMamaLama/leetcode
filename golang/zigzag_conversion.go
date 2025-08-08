package main

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	result := make([]string, numRows)
	isUp := true
	i := 0
	rowIdx := 0

	for i < len(s) {
		result[rowIdx] += string(s[i])
		if isUp {
			rowIdx++
		} else {
			rowIdx--
		}

		if rowIdx == numRows-1 {
			isUp = false
		}

		if rowIdx == 0 {
			isUp = true
		}

		i++
	}

	return strings.Join(result, "")
}
