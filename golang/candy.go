package main

func candy(ratings []int) int {
	var candies []int
	for i := 0; i < len(ratings); i++ {
		candies = append(candies, 1)
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	count := 0
	for _, val := range candies {
		count += val
	}

	return count
}
