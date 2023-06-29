package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	sum := maxArea(height)
	fmt.Println("Large container", sum)
}

func maxArea(height []int) int {
	right := len(height) - 1
	left := 0
	sum := 0
	for left <= right {
		temp := 0
		temp = max(height[left], height[right]) * (right - left)
		if sum < temp {
			sum = temp
		}

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return sum
}

func max(x int, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}
