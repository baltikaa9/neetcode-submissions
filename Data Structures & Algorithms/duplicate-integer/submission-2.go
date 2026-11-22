func hasDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for _, num := range nums {
		set[num] = struct{}{}
	}

	return len(set) != len(nums)
}