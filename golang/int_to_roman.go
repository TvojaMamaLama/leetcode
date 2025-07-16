package main

func intToRoman(num int) string {
	keySymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	keyValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""

	for num > 0 {
		for idx, value :=  range keyValues {
			if num >= value {
				result += keySymbols[idx]
				num -= value
				break
			}
		}
	}

	return result
}