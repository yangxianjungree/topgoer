package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		other := target - nums[i]
		if index, ok := m[other]; ok {
			return []int{index, i}
		}

		m[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println("two sum: ", twoSum(nums, 9))
}
