package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	pointer := 0
	for i := 0; i < len(t); i++ {
		if s[pointer] == t[i] {
			pointer++
			fmt.Println(i, pointer)
			if pointer == len(s) {
				return true
			}
		}
	}

	return false
}