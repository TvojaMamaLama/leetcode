package main

func isIsomorphic(s string, t string) bool {
	mapkaS := make(map[byte]byte)
	mapkaT := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if val, ok := mapkaS[s[i]]; ok && val != t[i] {
			return false
		}
		if val, ok := mapkaT[t[i]]; ok && val != s[i] {
			return false
		}
		mapkaS[s[i]] = t[i]
		mapkaT[t[i]] = s[i]
	}

	return true
}
