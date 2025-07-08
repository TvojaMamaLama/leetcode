package main

func twoSum2(numbers []int, target int) []int {
	history := make(map[int]int)

	for idx, val := range numbers {
		if idxVal, ok := history[target-val]; ok {
			return []int{idxVal, idx}
		}

		history[val] = idx
	}

	return []int{}
}

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < len(numbers); {
		current := numbers[i] + numbers[j]
		if current > target {
			j--
		} else if current < target {
			i++
		} else {
			return []int{i + 1, j + 1}
		}
	}

	return []int{1, 2}
}