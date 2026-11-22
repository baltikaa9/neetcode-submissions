package main

import "testing"

func Test(t *testing.T) {
	tests := []struct{
		name string
		s string
		expected bool
	}{
		{
			name: "test 1",
			s: "Was it a car or a cat I saw?",
			expected: true,
		},
		{
			name: "test 2",
			s: "tab a cat",
			expected: false,
		},
		{
			name: "signs",
			s: ".,!",
			expected: true,
		},
		{
			name: "empty",
			s: "",
			expected: true,
		},
		{
			name: "a",
			s: "a",
			expected: true,
		},
		{
			name: "test 3",
			s: "A man a plan a canal Panama",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			result := isPalindrome(tt.s)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}