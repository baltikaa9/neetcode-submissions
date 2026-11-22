package main

import (
	"slices"
	"testing"
)

func TestFindSubtreeSizes(t *testing.T) {
	tests := []struct {
		name     string
		parent   []int
		s        string
		expected []int
	}{
		{
			name:     "example 1",
			parent:   []int{-1, 0, 0, 1, 1, 1},
			s:        "abaabc",
			expected: []int{6, 3, 1, 1, 1, 1},
		},
		{
			name:     "example 2",
			parent:   []int{-1, 0, 4, 0, 1},
			s:        "abbba",
			expected: []int{5, 2, 1, 1, 1},
		},
		{
			name:     "example 3",
			parent:   []int{-1, 10, 0, 12, 10, 18, 11, 12, 2, 3, 2, 2, 2, 0, 4, 11, 4, 2, 0},
			s:        "babadabbdabcbaceeda",
			expected: []int{19, 1, 15, 2, 3, 1, 1, 1, 1, 1, 5, 2, 4, 1, 1, 1, 1, 1, 2},
		},
		{
			name:     "example 4",
			parent:   []int{-1, 2, 0},
			s:        "aab",
			expected: []int{3, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findSubtreeSizes(tt.parent, tt.s)

			if slices.Compare(result, tt.expected) != 0 {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
