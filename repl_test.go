package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Test cASE",
			expected: []string{"test", "case"},
		},
		{
			input:    "go PROGRAMMING",
			expected: []string{"go", "programming"},
		},
		{
			input:    "   hello   world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of the actual slice and the expected slice don't match")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("A word doesn't match an expected word")
			}
		}
	}
}
