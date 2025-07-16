package main

func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	for _, str := range strs {
		minLen = min(minLen, len(str))
	}

	result := ""
	for i := 0; i < minLen; i++ {
		current := strs[0][i]
		for _, str := range strs {
			if str[i] != current {
				return result
			}
		}

		result += string(current)
	}

	return result
}
