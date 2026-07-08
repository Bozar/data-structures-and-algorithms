// https://www.geeksforgeeks.org/dsa/find-element-appears-array-every-element-appears-twice/
package single

import ()

func Print(input []int) int {
	numToCount := map[int]bool{}
	for _, v := range input {
		if _, ok := numToCount[v]; ok {
			delete(numToCount, v)
		} else {
			numToCount[v] = true
		}
	}
	for k := range numToCount {
		return k
	}
	return 0
}
