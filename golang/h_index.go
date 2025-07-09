package main

import "slices"

func hIndex(citations []int) int {
	slices.SortFunc(citations, func(a, b int) int {
		return b - a
	})

	index := 0
	for i := 0; i < len(citations)-1; i++ {
		if citations[i] >= i+1 {
			index++
		} else {
			break
		}
	}

	return index
}
