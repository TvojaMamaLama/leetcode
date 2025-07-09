package main

import "fmt"

func removeDuplicates(nums []int) int {
	i, j := 1, 1
	current := nums[0]
	for j < len(nums) {
		if nums[j] != current {
			nums[i] = nums[j]
			current = nums[j]
			i++
		}
		j++
	}

	fmt.Println(nums)
	return i
}

func removeDuplicates2(nums []int) int {
	i, j := 1, 1
	current := nums[0]
	isFirst := true
	for j < len(nums) {
		fmt.Println(nums)
		fmt.Println(i, j, current, isFirst)
		if nums[j] != nums[j-1] {
			isFirst = true
		}

		if nums[j] == current && nums[j-1] == current && isFirst {
			isFirst = false
			i++
			j++
			continue
		}

		if nums[j] != current {
			nums[i] = nums[j]
			current = nums[j]
			i++
		}
		j++
	}

	return i
}