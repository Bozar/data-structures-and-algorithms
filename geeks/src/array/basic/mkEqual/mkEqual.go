// https://www.geeksforgeeks.org/dsa/minimum-increment-k-operations-make-elements-equal/
package mkEqual

import ()

func Print(input []int, k int) int {
	output := 0
	for _, v := range input {
		delta := v - input[0]
		if delta < 0 {
			delta *= -1
		}
		if delta%k != 0 {
			return -1
		}
		output += delta / k
	}
	return output
}
