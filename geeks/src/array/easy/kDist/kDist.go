// https://www.geeksforgeeks.org/dsa/check-given-array-contains-duplicate-elements-within-k-distance/
package kDist

import ()

func Print(input []int, k int) bool {
	valToIdx := map[int]int{}
	for i, v := range input {
		if _, ok := valToIdx[v]; !ok {
			valToIdx[v] = i
		} else if i-valToIdx[v] <= k {
			return true
		} else {
			valToIdx[v] = i
		}
	}
	return false
}
