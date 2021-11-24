package main

import "fmt"

func lengthOfLongestSubstring(input string) int {
	max := 1
	maxCur := 1
	m := make(map[byte]int, len(input))
	m[input[0]] = 0

	for i := 1; i < len(input); i++ {
		idx, ok := m[input[i]]
		if !ok || idx < i-maxCur {
			maxCur++
		} else {
			if max < maxCur {
				max = maxCur
			}
			maxCur = 1
		}
		m[input[i]] = i
	}
	if max < maxCur {
		max = maxCur
	}
	maxCur = 1

	return max
}

func main() {
	//
	input := "pwwkew"
	fmt.Println("Longest substring's length: ", lengthOfLongestSubstring(input))
}
