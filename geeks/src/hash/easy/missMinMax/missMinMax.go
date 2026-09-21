// https://www.geeksforgeeks.org/dsa/elements-to-be-added-so-that-all-elements-of-a-range-are-present-in-array/
package missMinMax

import ()

func Print(input []int) int {
	inputToBool := make(map[int]bool, len(input))
	minInput := input[0]
	maxInput := input[0]
	count := 0
	for _, v := range input {
		minInput = min(minInput, v)
		maxInput = max(maxInput, v)
		inputToBool[v] = true
	}
	for i := minInput; i < maxInput+1; i++ {
		if !inputToBool[i] {
			count++
		}
	}
	return count
}
