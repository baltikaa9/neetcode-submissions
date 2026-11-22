package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name string
		heights []int
		expected int
	}{
		{
			name: "test 1",
			heights: []int{7,1,7,2,2,4},
			expected: 8,
		},
		{
			name: "test 2",
			heights: []int{1,3,7},
			expected: 7,
		},
		{
			name: "test 3",
			heights: []int{3,1},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestRectangleArea(tt.heights)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestStack(t *testing.T) {
	tests := []struct {
		name string
		heights []int
		expected int
	}{
		{
			name: "test 1",
			heights: []int{7,1,7,2,2,4},
			expected: 8,
		},
		{
			name: "test 2",
			heights: []int{1,3,7},
			expected: 7,
		},
		{
			name: "test 3",
			heights: []int{3,1},
			expected: 3,
		},
		{
			name: "test 4",
			heights: []int{2,1,5,6,2,3},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestRectangleAreaStack(tt.heights)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}