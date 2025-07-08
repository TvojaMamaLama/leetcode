package main

func rotate(nums []int, k int) {
	for i := 0; i+k < len(nums); i++ {
		k %= len(nums)
	}
}
