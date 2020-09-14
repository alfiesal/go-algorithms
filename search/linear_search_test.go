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
	{[]string{"A", "A", "B", "B", "C", "C", "D", "E"}, "D", 6, "doubled"},
	{[]string{"A", "B", "C", "D", "F"}, "A", 0, "first-position"},
	{[]string{"A", "B", "C", "D", "F"}, "F", 4, "last-position"},
}

func TestLinearSearch(t *testing.T) {
	for _, test := range searchTests {
		actual := linearSearch(test.input, test.query)
		if actual != test.expected {
			t.Errorf("Test %s failed: %v inputed, %d excepted, %d received", test.name, test.input, test.expected, actual)
		}
	}
}
