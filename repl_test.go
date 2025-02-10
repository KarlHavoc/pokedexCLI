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
			input:    " hello  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HOW manY pokemon are there",
			expected: []string{"how", "many", "pokemon", "are", "there"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("actual: %v - expected: %v", actual, c.expected)
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual word: %s - expected word: %s", word, expectedWord)
				t.Fail()
			}
		}
	}
}
