package search

import (
	"math"
)

func jumpSearch(input []int, query int) int {
	jumpSize := int(math.Sqrt(float64(len(input))))

	index := 0

	for index < len(input) {
		if query == input[index] {
			return index
		}

		if query > input[index] {
			index = index + jumpSize
		}

		if query < input[index] {
			max := index
			index = index - jumpSize

			for index < max {
				if query == input[index] {
					return index
				}
				index = index + 1
			}

			return -1
		}
	}

	return -1
}
