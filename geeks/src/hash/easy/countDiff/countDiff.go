// https://www.geeksforgeeks.org/dsa/count-pairs-difference-equal-k/
package countDiff

import ()

func Print(input []int, target int) int {
	inputToCount := make(map[int]int, len(input))
	count := 0
	for _, v := range input {
		count += inputToCount[v-target]
		count += inputToCount[v+target]
		inputToCount[v] += 1
	}
	return count
}
