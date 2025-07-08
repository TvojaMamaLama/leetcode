package main

import "unicode/utf8"

func strStr(haystack string, needle string) int {
	needleLen := utf8.RuneCountInString(needle)
	if needleLen > utf8.RuneCountInString(haystack) {
		return -1
	}

	for i := 0; i < utf8.RuneCountInString(haystack)-needleLen+1; i++ {
		if haystack[i] == needle[0] {
			if haystack[i:i+needleLen] == needle {
				return i
			}
		}
	}

	return -1
}