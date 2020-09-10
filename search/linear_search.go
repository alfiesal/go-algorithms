package main

import "fmt"

func linearSearch(input []string, query string) int {

	for index, item := range input {
		if item == query {
			return index
		}
	}
	return -1
}

func main() {
	fmt.Println("####################################")
	fmt.Println("Linear Search - complexity O(n)")
	fmt.Println("####################################")

	input := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	query := "J"
	index := linearSearch(input, query)

	if index == -1 {
		fmt.Println("Can't find element: ", query)

		return
	}

	fmt.Println("Index of the found element: ", index)
	fmt.Println("Value of found element: ", input[index])
}
