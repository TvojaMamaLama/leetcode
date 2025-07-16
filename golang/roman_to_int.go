package main

func romanToInt(s string) int {
	d := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}
	result := 0

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			result += d[s[i]]
			break
		}

		if d[s[i]] >= d[s[i+1]] {
			result += d[s[i]]
		} else {
			result += d[s[i+1]] - d[s[i]]
			i++
		}
	}

	return result
}
