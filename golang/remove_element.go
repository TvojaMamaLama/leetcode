package main

func removeElement(nums []int, val int) int {
	f, l := 0, len(nums)-1
	c := 0

	for f <= l {
		if nums[l] == val {
			l--
			c++
			continue
		}

		if nums[f] == val {
			nums[l], nums[f] = nums[f], nums[l]
			l--
			f++
			c++
			continue
		}

		f++
	}

	return len(nums) - c
}