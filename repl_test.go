package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Coded in Go    ",
			expected: []string{"coded", "in", "go"},
		},
		{
			input:    "  Everyone should watch Superman  ",
			expected: []string{"everyone", "should", "watch", "superman"},
		},
		// add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match, actual = %d; expected %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if len(word) != len(expectedWord) {
				t.Errorf("lengths don't match, word = %d %s; expectedWord %d %s", len(word), word, len(expectedWord), expectedWord)
			}
		}
	}
}
