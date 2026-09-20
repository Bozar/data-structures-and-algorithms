// https://www.geeksforgeeks.org/dsa/count-pairs-with-given-sum/
package countSum

import ()

func Print(input []int, target int) int {
	inputToCount := make(map[int]int, len(input))
	count := 0
	for _, v := range input {
		count += inputToCount[target-v]
		inputToCount[v] += 1
	}
	return count
}
