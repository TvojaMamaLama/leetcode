package main

func lengthOfLongestSubstring(s string) int {
	mapa := make(map[byte]int)
	start := 0
	best := 0

	for i := 0; i < len(s); i++ {
		if idx, ok := mapa[s[i]]; ok && idx >= start {
			best = max(best, i-start)
			start = idx + 1
		}

		mapa[s[i]] = i
	}

	best = max(best, len(s)-start)

	return best
}