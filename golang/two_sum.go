package main

func twoSum(nums []int, target int) []int {
	mapka := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if val, ok := mapka[target-nums[i]]; ok {
			return []int{val, i}
		}

		mapka[nums[i]] = i
	}

	return []int{}
}
