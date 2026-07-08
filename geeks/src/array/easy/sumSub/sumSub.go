// https://www.geeksforgeeks.org/dsa/sum-of-all-subarrays/
package sumSub

import ()

func Print(input []int) int {
	sum := 0
	lenIn := len(input)
	for i, v := range input {
		sum += v * (i + 1) * (lenIn - i)
	}
	return sum
}
