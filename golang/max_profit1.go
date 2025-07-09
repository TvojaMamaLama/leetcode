package main

func maxProfit1(prices []int) int {
	best := 0
	lIdx := 0
	hIdx := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[lIdx] {
			lIdx = i
		} else if prices[i] > prices[hIdx] || lIdx > hIdx {
			hIdx = i
		}

		if lIdx < hIdx {
			best = max(best, prices[hIdx]-prices[lIdx])
		}
	}

	if lIdx < hIdx {
		best = max(best, prices[hIdx]-prices[lIdx])
	}

	return best
}