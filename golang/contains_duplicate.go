package main

func containsNearbyDuplicate(nums []int, k int) bool {
	mapka := make(map[int]int)

	for idx, num := range nums {
		if val, ok := mapka[num]; ok {
			if idx-val <= k {
				return true
			}
		}
		mapka[num] = idx
	}

	return false
}
