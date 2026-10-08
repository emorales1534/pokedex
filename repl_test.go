package main

import (
	"testing"
)

func TestCleanInput (t *testing.T) {
	cases := []struct {
		input	string
		expected []string}{
			{
				input: " hello world",
				expected: []string{"hello","world"},
			},
		}
	
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("different lengths")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word %v does not match %v", word, expectedWord)
			}
		}
	}
}
