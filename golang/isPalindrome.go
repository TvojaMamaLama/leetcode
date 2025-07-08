package main

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	p1 := 0
	p2 := utf8.RuneCountInString(s) - 1
	for p1 < p2 {
		for p1 <= p2 {
			if unicode.IsLetter(rune(s[p1])) || unicode.IsNumber(rune(s[p1])) {
				break
			}
			p1++
		}

		for p1 <= p2 {
			if unicode.IsLetter(rune(s[p2])) || unicode.IsNumber(rune(s[p2])) {
				break
			}
			p2--
		}

		if s[p1] == s[p2] {
			p1++
			p2--
		} else {
			return false
		}
	}

	return true
}