//https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation

package main

import "neetcode-submissions/tests/matrix"

func rotate(mat [][]int) {
	n := len(mat)

	for i := range n {
		for j := i + 1; j < n; j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}

	for i := range n {
		for j := range n / 2 {
			mat[i][j], mat[i][n-1-j] = mat[i][n-1-j], mat[i][j]
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func findRotation(mat [][]int, target [][]int) bool {
	if equal(mat, target) {
		return true
	}

	for range 3 {
		rotate(mat)
		if equal(mat, target) {
			return true
		}
	}

	return false
}

func main() {
	mat := matrix.MatrixInt{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	target := matrix.MatrixInt{
		{1, 1, 1},
		{0, 1, 0},
		{0, 0, 0},
	}

	println(findRotation(mat, target))
}
