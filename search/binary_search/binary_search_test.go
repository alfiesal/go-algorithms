package binary_search

import "testing"

type binarySearchStruct struct {
	input    []int
	query    int
	expected int
	name     string
}

var binarySearchData = []binarySearchStruct{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 8, "default"},
	{[]int{1, 3, 4, 5, 6, 7, 9, 10}, 2, -1, "absent"},
	{[]int{}, 2, -1, "empty"},
	{[]int{1, 1, 2, 2, 2, 3, 3, 4, 5, 8, 9}, 8, 9, "doubled"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, "first-position"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, "last-position"},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range binarySearchData {
		actual := iterativeBinarySearch(test.input, test.query)
		if actual != test.expected {
			t.Errorf("Iterative test %s failed: %v inputed, %d excepted, %d received", test.name, test.input, test.expected, actual)
		}
	}

	for _, test := range binarySearchData {
		actual := binarySearch(test.input, test.query, 0, len(test.input)-1)
		if actual != test.expected {
			t.Errorf("Recursive test %s failed: %v inputed, %d excepted, %d received", test.name, test.input, test.expected, actual)
		}
	}
}
