package search

import "fmt"

func binarySearch(input []int, query int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}

	middleIndex := int(lowIndex + (highIndex-lowIndex)/2)

	if query == input[middleIndex] {
		return middleIndex
	}

	if query < input[middleIndex] {
		return binarySearch(input, query, lowIndex, middleIndex-1)
	}

	if query > input[middleIndex] {
		return binarySearch(input, query, middleIndex+1, highIndex)
	}

	return -1
}

// Example of usage binary search algorithm
func BinarySearchExample() {
	fmt.Println("####################################")
	fmt.Println("Binary Search - complexity O(log2n)")
	fmt.Println("####################################")

	input := []int{1, 2, 3, 4, 8, 10, 11, 20, 21, 24, 29, 30}
	query := 11
	index := binarySearch(input, query, 0, len(input))

	if index == -1 {
		fmt.Println("Can't find element: ", query)

		return
	}

	fmt.Println("Index of the found element: ", index)
	fmt.Println("Value of found element: ", input[index])
}
