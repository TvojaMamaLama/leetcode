package main

import (
	"strconv"
)

func isHappy(n int) bool {
	mapka := make(map[int]struct{})
	mapka[n] = struct{}{}

	for {
		s := strconv.Itoa(n)
		sum := 0
		for _, v := range s {
			intV, _ := strconv.Atoi(string(v))
			sum += intV * intV
		}

		if sum == 1 {
			return true
		}

		if _, ok := mapka[sum]; ok {
			return false
		}

		mapka[sum] = struct{}{}
		n = sum
	}
}
