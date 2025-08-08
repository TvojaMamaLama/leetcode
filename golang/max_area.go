package main

import "fmt"

func maxArea(height []int) int {
	area := 0

	for i, j := 0, len(height)-1; i < j; {
		area = max(min(height[i], height[j])*(j-i), area)
		fmt.Println(height[i], height[j])
		if height[i] < height[j] {
			i++
		} else {
			j--
		}

		fmt.Println(i, j, area)
	}

	return area
}
