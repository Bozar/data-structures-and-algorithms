// https://www.geeksforgeeks.org/dsa/complete-guide-on-array-rotations/
package rotate

import ()

func Print(input []int, distance int) []int {
	lenIn := len(input)
	distance %= lenIn
	left := make([]int, lenIn-distance)
	right := make([]int, distance)
	lenLe := len(left)
	for i, v := range input {
		if i < lenLe {
			left[i] = v
		} else {
			right[i-lenLe] = v
		}
	}
	right = append(right, left...)
	return right
}
