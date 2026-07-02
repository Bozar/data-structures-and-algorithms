// https://www.geeksforgeeks.org/dsa/remove-duplicates-sorted-array/
package rmvDup

import ()

func Print(input []int) []int {
	keep := 1
	for i := 1; i < len(input); i++ {
		if input[keep-1] != input[i] {
			input[keep] = input[i]
			keep++
		}
	}
	return input[0:keep]
}
