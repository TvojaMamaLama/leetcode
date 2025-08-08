package main

func minSubArrayLen(target int, nums []int) int {
	i, sum := 0, 0          // i = 2
	minLen := len(nums) + 1 // 4
	// 2,3,1,!2,4,3
	// 7
	for j := 0; j < len(nums); j++ { // j = 5
		sum += nums[i] // sum = 10
		for sum >= target {
			minLen = min(minLen, j-i+1)
			sum -= nums[j]
			i++
		}
	}

	if minLen > len(nums) {
		return 0
	}

	return minLen
}
