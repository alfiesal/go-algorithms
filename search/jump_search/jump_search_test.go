package jump_search

import "testing"

type jumpSearchStruct struct {
	input    []int
	query    int
	expected int
	name     string
}

var jumpSearchData = []jumpSearchStruct{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 8, "default"},
	{[]int{1, 3, 4, 5, 6, 7, 9, 10}, 2, -1, "absent"},
	{[]int{}, 2, -1, "empty"},
	{[]int{1, 1, 2, 2, 2, 3, 3, 4, 5, 8, 9}, 8, 9, "doubled"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, "first-position"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, "last-position"},
}

func TestJumpSearch(t *testing.T) {
	for _, test := range jumpSearchData {
		actual := jumpSearch(test.input, test.query)
		if actual != test.expected {
			t.Errorf("Test %s failed: %v inputed, %d excepted, %d received", test.name, test.input, test.expected, actual)
		}
	}
}
