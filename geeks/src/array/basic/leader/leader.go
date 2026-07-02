// https://www.geeksforgeeks.org/dsa/leaders-in-an-array/
package leader

import ()

func Print(input []int) []int {
	maxVal := -1
	output := make([]int, 0, len(input))
	for i := len(input) - 1; i > -1; i-- {
		if input[i] > maxVal {
			output = append(output, input[i])
			maxVal = input[i]
		}
	}
	for i, j := 0, len(output)-1; i < j; {
		output[i], output[j] = output[j], output[i]
		i++
		j--
	}
	return output
}
