package main

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	result := make([]string, 0, len(words))

	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			result = append(result, words[i])
		}
	}
	
	return strings.Join(result, " ")
}
