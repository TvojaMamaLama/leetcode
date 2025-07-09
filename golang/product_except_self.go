package main

func productExceptSelf(nums []int) []int {
	left, right := make([]int, len(nums)), make([]int, len(nums))

	left[0], right[len(right)-1] = 1, 1

	for i := 1; i < len(left); i++ {
		left[i] = nums[i-1] * left[i-1]
	}

	for i := len(right) - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	result := make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		result[i] = left[i] * right[i]
	}

	return result
}

func productExceptSelfBest(nums []int) []int {
	n := len(nums)
	right := make([]int, n)
	res := make([]int, n)

	prod := 1
	for i := n - 1; i >= 0; i-- {
		prod *= nums[i]
		right[i] = prod
	}

	prod = 1
	for i := 0; i < n-1; i++ {
		lp := prod
		rp := right[i+1]

		res[i] = rp * lp
		prod *= nums[i]
	}
	res[n-1] = prod
	return res
}
