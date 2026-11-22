// https://leetcode.com/problems/flip-square-submatrix-vertically/

package main

import "neetcode-submissions/tests/matrix"

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	for i, j := x, x+k-1; i < j; i, j = i+1, j-1 {
		for n := y; n < y+k; n++ {
			grid[i][n], grid[j][n] = grid[j][n], grid[i][n]
		}
	}

	return grid
}

func main() {
	grid := matrix.MatrixInt{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	matrix.MatrixInt(reverseSubmatrix(grid, 1, 0, 3)).Print()
}
