// https://www.geeksforgeeks.org/dsa/rearrange-array-such-that-even-positioned-are-greater-than-odd/
package evGreater

import ()

func Print(input []int) []int {
	for i := range input {
		if i == 0 {
			continue
		} else if (i%2 == 0) && (input[i] > input[i-1]) {
			input[i-1], input[i] = input[i], input[i-1]
		} else if (i%2 != 0) && (input[i] < input[i-1]) {
			input[i-1], input[i] = input[i], input[i-1]
		}
	}
	return input
}
