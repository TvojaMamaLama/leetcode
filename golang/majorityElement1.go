package main

func majorityElement1(nums []int) int {
	counter := make(map[int]int)
	best := 0

	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++

		if counter[i] > counter[best] {
			best = i
		}
	}

	return best
}