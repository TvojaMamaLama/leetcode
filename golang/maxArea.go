package main

func maxArea(height []int) int {
	area := 0
	for i, j := 0, len(height)-1; i < j; {
		area = max(area, min(height[i], height[j])*(j-i))
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}

	return area
}