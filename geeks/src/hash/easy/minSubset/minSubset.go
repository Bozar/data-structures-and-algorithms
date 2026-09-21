// https://www.geeksforgeeks.org/dsa/minimum-number-subsets-distinct-elements/
package minSubset

import ()

func Print(input []int) int {
	inputToInt := make(map[int]int, len(input))
	count := 0
	for _, v := range input {
		inputToInt[v] += 1
	}
	for _, v := range inputToInt {
		count = max(count, v)
	}
	return count
}
