// https://www.geeksforgeeks.org/dsa/remove-minimum-number-elements-no-common-element-exist-array/
package noCommon

import ()

func Print(input1 []int, input2 []int) int {
	inputToCount := make(map[int]int, len(input1))
	count := 0
	for _, v := range input1 {
		inputToCount[v] += 1
	}
	for _, v := range input2 {
		if inputToCount[v] > 0 {
			inputToCount[v] -= 1
			count += 1
		}
	}
	return count
}
