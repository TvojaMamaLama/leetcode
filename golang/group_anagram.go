package main

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	mapka := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		sortedStr := string(runes)
		mapka[sortedStr] = append(mapka[sortedStr], str)
	}
	var result [][]string
	for _, str := range mapka {
		result = append(result, str)
	}

	return result
}
