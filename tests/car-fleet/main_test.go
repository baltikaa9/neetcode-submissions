package main

import "testing"

func TestCarFleet(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		position []int
		speed    []int
		expected int
	}{
		{
			name:     "example 1 - merge at destination",
			target:   10,
			position: []int{1, 4},
			speed:    []int{3, 2},
			expected: 1,
		},
		{
			name:     "example 2 - three fleets",
			target:   10,
			position: []int{4, 1, 0, 7},
			speed:    []int{2, 2, 1, 1},
			expected: 3,
		},
		{
			name:     "middle car already merged into fleet",
			target:   10,
			position: []int{0, 5, 7},
			speed:    []int{3, 2, 1}, // times: [10/3, 1.0, 3.0]
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := carFleet(tt.target, tt.position, tt.speed)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
