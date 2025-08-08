package main

func canConstruct(ransomNote string, magazine string) bool {
	mapkaNote := map[rune]int{}
	mapkaMagazine := map[rune]int{}

	for _, r := range ransomNote {
		mapkaNote[r]++
	}

	for _, r := range magazine {
		mapkaMagazine[r]++
	}

	for key, value := range mapkaNote {
		if valueMag, ok := mapkaMagazine[key]; ok {
			if valueMag < value {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
