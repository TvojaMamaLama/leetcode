package main

import "strings"

func wordPattern(pattern string, s string) bool {
	mapka := make(map[byte]string)
	words := strings.Split(s, " ")

	for i := 0; i < len(pattern); i++ {
		if val, ok := mapka[pattern[i]]; !ok {
			mapka[pattern[i]] = words[i]
		} else {
			if val != words[i] {
				return false
			}
		}
	}

	return true
}
