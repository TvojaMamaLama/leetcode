package main

func jump(nums []int) int {
	maxReachable := 0
	currentEnd := 0
	jumps := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] >= maxReachable {
			maxReachable = i + nums[i]
		}

		if i >= currentEnd {
			currentEnd = maxReachable
			jumps++
		}
	}

	return jumps
}
