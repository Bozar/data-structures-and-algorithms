// https://www.geeksforgeeks.org/dsa/move-zeroes-end-array/
package mvZero

import ()

func Print(input []int) []int {
	slow := 0
	for i, v := range input {
		if v == 0 {
			continue
		}
		input[slow], input[i] = input[i], input[slow]
		slow += 1
	}
	return input
}
