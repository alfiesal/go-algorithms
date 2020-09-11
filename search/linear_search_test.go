package search

import "testing"

type searchTest struct {
	input    []string
	query    string
	expected int
	name     string
}

var searchTests = []searchTest{
	{[]string{"A", "B", "C", "D", "F"}, "B", 1, "default"},
	{[]string{"A", "C", "D"}, "B", -1, "absent"},
	{[]string{}, "B", -1, "empty"},
}

func TestLinearSearch(t *testing.T) {
	for _, test := range searchTests {
		actual := linearSearch(test.input, test.query)
		if actual != test.expected {
			t.Error("Test failed: {} inputed, {} excepted, {} received", test.input, test.expected, actual)
		}
	}
}
