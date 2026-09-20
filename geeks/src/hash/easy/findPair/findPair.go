// https://www.geeksforgeeks.org/dsa/check-if-pair-with-given-sum-exists-in-array/
package findPair

import ()

func Print(input []int, target int) bool {
	inputToBool := make(map[int]bool, len(input))
	for _, v := range input {
		if inputToBool[v] {
			return true
		}
		inputToBool[target-v] = true
	}
	return false
}
