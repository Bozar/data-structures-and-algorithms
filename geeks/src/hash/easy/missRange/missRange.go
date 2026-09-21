// https://www.geeksforgeeks.org/dsa/find-missing-elements-of-a-range/
package missRange

import ()

func Print(input []int, low int, high int) []int {
	inputToBool := make(map[int]bool, high-low+1)
	output := make([]int, 0, high-low+1)
	for _, v := range input {
		if (v >= low) && (v <= high) {
			inputToBool[v] = true
		}
	}
	for i := low; i < high+1; i++ {
		if !inputToBool[i] {
			output = append(output, i)
		}
	}
	return output
}
