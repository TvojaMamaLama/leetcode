package main

func candy(ratings []int) int {
	candies := make([]int, len(ratings))

	for i := 0; i < len(candies); i++ {
		candies[i] = 1
	}

	for i := 1; i < len(candies); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(candies) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = max(candies[i], candies[i+1]) + 1
		}
	}

	result := 0
	for _, val := range candies {
		result += val
	}

	return result
}
