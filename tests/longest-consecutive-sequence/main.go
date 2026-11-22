// https://neetcode.io/problems/longest-consecutive-sequence/

package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]int, len(nums))

	for _, n := range nums {
		m[n] = 1
	}
	
	longest := 0
	
	for num := range m {
		if _, exists := m[num-1]; exists {
			continue
		}
		
		curNum := num
		
		for {
			if _, exists := m[curNum + 1]; exists {
				m[num]++
				curNum++
			} else {
				break
			}
		}
		
		if m[num] > longest {
			longest = m[num]
		}
	}

	return longest
}

func main() {
	nums := []int{2, 20, 4, 10, 3, 4, 5}

	fmt.Println(longestConsecutive(nums))

	nums = []int{0, 3, 2, 5, 4, 6, 1, 1}

	fmt.Println(longestConsecutive(nums))
}
