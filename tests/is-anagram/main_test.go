package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "example 1 - anagram",
			s:        "racecar",
			t:        "carrace",
			expected: true,
		},
		{
			name:     "example 2 - non anagram",
			s:        "jar",
			t:        "jam",
			expected: false,
		},
		// {
		// 	name:     "example 3 - ru",
		// 	s:        "хам",
		// 	t:        "мах",
		// 	expected: true,
		// },
		{
			name:     "example 3 - empty",
			s:        "",
			t:        "",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.t, tt.s)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
