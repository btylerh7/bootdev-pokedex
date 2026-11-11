package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{

			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " what is this ",
			expected: []string{"what", "is", "this"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Logf("Expected: %s; Actual: %s;", c.expected, actual)
			t.Errorf("Lengths different. Expected legnth of %d, got %d", len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {

			}
		}

	}

}
