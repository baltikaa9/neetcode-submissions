// https://neetcode.io/problems/duplicate-integer/

package main

func hasDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for _, num := range nums {
		set[num] = struct{}{}
	}

	return len(set) != len(nums)
}

func main() {
	nums := []int{1, 2, 3, 4}

	println(hasDuplicate(nums))
}