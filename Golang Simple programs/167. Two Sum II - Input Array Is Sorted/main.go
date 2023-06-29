package main

func twoSum(numbers []int, target int) []int {
	right := len(numbers) - 1
	left := 0

	for left <= right {
		sum := numbers[right] + numbers[left]
		if target < sum {
			right--
		} else if target > sum {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return []int{}
}
