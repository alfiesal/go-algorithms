package native_search

import "fmt"

func nativeSearch(text string, pattern string) []int {
	var positions []int

	return positions
}

func NativeSearchExample() {
	fmt.Println("#########################################")
	fmt.Println("Native String Search - complexity O(n*m)")
	fmt.Println("n = length of text, m = length of pattern")
	fmt.Println("#########################################")

	input := "ACADGABAABCAADABCDS"
	query := "ABCD"

	position := nativeSearch(input, query)

	if len(position) == 0 {
		fmt.Println("Pattern not found in given text!")

		return
	}
}
