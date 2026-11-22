// https://neetcode.io/problems/daily-temperatures/

package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))

	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}

func main() {
	temps := []int{30, 38, 30, 36, 35, 40, 28}

	fmt.Println(dailyTemperatures(temps))
}
