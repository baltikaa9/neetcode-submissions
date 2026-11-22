// https://leetcode.com/problems/find-subtree-sizes-after-changes/

package main

func dfs(node int, parent [][]int, newParent, answer []int, s string, letters *[26][]int) int {
	letters[s[node]-'a'] = append(letters[s[node]-'a'], node)

	for _, i := range parent[node] {
		l := s[i] - 'a'

		if len(letters[l]) > 0 {
			newParent[i] = letters[l][len(letters[l])-1]
		}

		answer[newParent[i]] += dfs(i, parent, newParent, answer, s, letters)
		letters[l] = letters[l][:len(letters[l])-1]
	}

	return answer[node]
}

func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)

	adjacency := make([][]int, n, n)

	for i := range n {
		p := parent[i]

		if p == -1 {
			continue
		}

		adjacency[p] = append(adjacency[p], i)
	}

	newParent := make([]int, n, n)
	copy(newParent, parent)

	answer := make([]int, n, n)

	for i := range answer {
		answer[i] = 1
	}

	letters := [26][]int{}

	dfs(0, adjacency, newParent, answer, s, &letters)

	return answer
}

func main() {
	parent := []int{-1, 0, 0, 1, 1, 1}
	s := "abaabc"

	// parent := []int{-1, 0, 4, 0, 1}
	// s := "abbba"

	println(findSubtreeSizes(parent, s))
}
