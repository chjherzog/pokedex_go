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
			input:    "HellO    World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORlD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello this is my world",
			expected: []string{"hello", "this", "is", "my", "world"},
		},
		{
			input:    "   hello    world    ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d words but got %d words", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if word != expected {
				t.Errorf("Expected %v but got %v", word, expected)
				break
			}
		}
	}

}
