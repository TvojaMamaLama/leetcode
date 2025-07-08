package main

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}

	wLen := len(words[0])
	wCount := len(words)
	result := []int{}

	wFreq := make(map[string]int)
	for _, word := range words {
		wFreq[word]++
	}

	for i := 0; i < wCount; i++ {
		left, right := i, i
		seenWords := make(map[string]int)
		count := 0

		for right+wLen < len(s) {
			word := s[right : right+wLen]
			right += wLen

			if _, ok := wFreq[word]; ok {
				seenWords[word]++
				count++

				for seenWords[word] > wFreq[word] {
					leftWord := s[left : left+wLen]
					seenWords[leftWord]--
					count--
					left += wLen
				}

				if count == wCount {
					result = append(result, left)
				}
			} else {
				seenWords = make(map[string]int)
				count = 0
				left = right
			}
		}
	}

	return result
}